package main

import (
	"container/list"
	"errors"
	"runtime"
	"sync"
	"time"
)

type Cache struct {
	*cache
}

type cache struct {
	values map[interface{}]*cacheEntry
	freqs  *list.List
	// If len > maxLen, cache will automatically evict down to MaxLen
	len           int
	maxLen        int
	cleanInterval time.Duration
	janitor       *janitor
	lock          *sync.Mutex
	name          string
	enableMetrics bool
	total         int // total request
}

type cacheEntry struct {
	key        interface{}
	value      interface{}
	expiresAt  *time.Time
	lastAccess time.Time
	freqNode   *list.Element
	hit        int
}

type listEntry struct {
	entries map[*cacheEntry]byte
	freq    int
}

// maxLen: cache max size
// ci: delete keys that have expired or not accessed for long time
func newCache(maxLen int, ci time.Duration, name string) *cache {
	c := new(cache)
	c.values = make(map[interface{}]*cacheEntry)
	c.freqs = list.New()
	c.lock = new(sync.Mutex)
	c.maxLen = maxLen
	c.name = name
	c.cleanInterval = ci
	c.enableMetrics = len(name) != 0
	return c
}

func (c *cache) Get(key interface{}, threshold float64) (interface{}, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.total++
	if e, ok := c.values[key]; ok {
		if e.expiresAt != nil {
			if e.expiresAt.After(time.Now()) {
				return c.value(e, threshold)
			}
			c.del(e)
			return nil, nil
		}
		return c.value(e, threshold)
	}
	return nil, nil
}

func (c *cache) value(e *cacheEntry, threshold float64) (interface{}, error) {
	c.increment(e)
	// check threshold
	if e.hit <= 1 || float64(e.hit) < float64(c.total)*threshold {
		return nil, errors.New("Not hit threshold")
	}
	return e.value, nil
}

func (c *cache) Set(key interface{}, value interface{}, expiresAt *time.Time) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.values[key]; ok {
		// value already exists for key.  overwrite
		e.value = value
		e.expiresAt = expiresAt
		c.increment(e)
	} else {
		// value doesn't exist.  insert
		e := new(cacheEntry)
		e.key = key
		e.value = value
		e.expiresAt = expiresAt
		c.values[key] = e
		c.increment(e)
		c.len++
		if c.len > c.maxLen*2 {
			c.evict(c.len - c.maxLen)
		}
	}
}

func (c *cache) Remove(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if entry, ok := c.values[key]; ok {
		c.del(entry)
	}
}

func (c *cache) Len() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.len
}

func (c *cache) Evict(count int) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.evict(count)
}

func (c *cache) evict(count int) int {
	// No lock here so it can be called
	// from within the lock (during Set)
	var evicted int
	for i := 0; i < count; {
		if place := c.freqs.Front(); place != nil {
			for entry := range place.Value.(*listEntry).entries {
				if i < count {
					c.del(entry)
					evicted++
					i++
				}
			}
		}
	}
	return evicted
}

func (c *cache) increment(e *cacheEntry) {
	currentPlace := e.freqNode
	var nextFreq int
	var nextPlace *list.Element
	if currentPlace == nil {
		// new entry
		nextFreq = 1
		nextPlace = c.freqs.Front()
	} else {
		// move up
		nextFreq = currentPlace.Value.(*listEntry).freq + 1
		nextPlace = currentPlace.Next()
	}

	if nextPlace == nil || nextPlace.Value.(*listEntry).freq != nextFreq {
		// create a new list entry
		li := new(listEntry)
		li.freq = nextFreq
		li.entries = make(map[*cacheEntry]byte)
		if currentPlace != nil {
			nextPlace = c.freqs.InsertAfter(li, currentPlace)
		} else {
			nextPlace = c.freqs.PushFront(li)
		}
	}
	e.freqNode = nextPlace
	e.lastAccess = time.Now()
	e.hit++
	nextPlace.Value.(*listEntry).entries[e] = 1
	if currentPlace != nil {
		// remove from current position
		c.remEntry(currentPlace, e)
	}
}

func (c *cache) remEntry(place *list.Element, entry *cacheEntry) {
	entries := place.Value.(*listEntry).entries
	delete(entries, entry)
	if len(entries) == 0 {
		c.freqs.Remove(place)
	}
}

func (c *cache) del(entry *cacheEntry) {
	delete(c.values, entry.key)
	c.remEntry(entry.freqNode, entry)
	c.len--
}

func (c *cache) Purge() {
	c.lock.Lock()
	defer c.lock.Unlock()

	purgeTime := time.Now().Add(-c.cleanInterval)
	for _, e := range c.values {
		e.hit = 0 // reset hit
		if e.expiresAt != nil && e.expiresAt.Before(time.Now()) ||
			e.lastAccess.Before(purgeTime) {
			c.del(e)
		}
	}

	c.total = 0 // reset total
}

func (c *cache) MGet(keys []interface{}, threshold float64) (map[interface{}]interface{}, map[interface{}]error) {
	r := make(map[interface{}]interface{})
	e := make(map[interface{}]error)
	for _, k := range keys {
		r[k], e[k] = c.Get(k, threshold)
	}

	return r, e
}

func (c *cache) MSet(kvs map[interface{}]interface{}, expiresAt *time.Time) {
	for k, v := range kvs {
		c.Set(k, v, expiresAt)
	}
}

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			c.Purge()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j
	go j.Run(c)
}

func New(maxLen int, ci time.Duration, name string) *Cache {
	c := newCache(maxLen, ci, name)
	C := &Cache{c}
	if ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}
