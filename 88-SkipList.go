package main

import (
	"math/rand"
	"fmt"
	"strings"
	"strconv"
	"testing"
)

type Value string
type Key int
type SkipList struct {
	head skiplistNode //跳表的头结点
	size int          //跳表包含的元素个数
}

//跳表结点定义
type skiplistNode struct {
	level []*skiplistNode //每个结点包含多个后向指针
	value Value           //结点的值
	key   Key             //结点的key
}

//跳表构造函数
func NewSkipList() *SkipList {
	sl := &SkipList{}
	sl.head.level = make([]*skiplistNode, 1)
	sl.head.value = "root"
	sl.size = 1
	return sl
}

//查询操作
func (sl *SkipList) Get(key Key) *Value {
	now := &sl.head
	layer := len(sl.head.level) - 1
	for layer >= 0 {
		nex := now.level[layer]
		if nex == nil {
			layer--
		} else if nex.key > key {
			layer --
		} else if nex.key == key {
			break
		} else {
			now = nex
		}
	}
	if layer < 0 {
		return nil
	} else {
		return &now.value
	}
}

//删除操作
func (sl *SkipList) Del(key Key) {
	sl.size--
	now := &sl.head
	layer := len(sl.head.level) - 1
	for layer >= 0 {
		nex := now.level[layer]
		if nex == nil {
			layer--
		} else if nex.key > key {
			layer --
		} else if nex.key == key {
			break
		} else {
			now = nex
		}
	}
	//元素本来就不存在
	if layer < 0 {
		return
	}
	nex := now.level[layer]
	if nex == nil {
		return
	}
	for i := 0; i < len(nex.level); i++ {
		now.level[i] = nex.level[i]
	}
	sl.cutHeight()
}

//削减root的高度
func (sl *SkipList) cutHeight() {
	totalLayer := len(sl.head.level)
	for i := totalLayer - 1; i >= 0; i-- {
		if sl.head.level[i] == nil {
			totalLayer --
		} else {
			break
		}
	}
	if totalLayer <= 0 {
		totalLayer = 1
	}
	sl.head.level = sl.head.level[:totalLayer]
}
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

//对于新插入的结点随机它的高度
func randomHeight() int {
	layerCount := 1
	for {
		newLayer := rand.Float32()
		if newLayer > 0.5 {
			layerCount++
		} else {
			break
		}
	}
	return layerCount
}

//增加root的高度
func (sl *SkipList) growHeight(layerCount int) {
	if layerCount < len(sl.head.level) {
		return
	}
	temp := sl.head.level
	sl.head.level = make([]*skiplistNode, layerCount)
	for i := 0; i < len(temp); i++ {
		sl.head.level[i] = temp[i]
	}
}

//添加元素操作
func (sl *SkipList) Put(key Key, value Value) {
	sl.size++
	layerCount := randomHeight()
	if layerCount > len(sl.head.level) {
		sl.growHeight(layerCount)
	}
	no := skiplistNode{level: make([]*skiplistNode, layerCount), key: key, value: value}
	now := &sl.head
	layer := len(sl.head.level) - 1
	lastNode := make([]*skiplistNode, layerCount)
	for i := 0; i < layerCount; i++ {
		lastNode[i] = &sl.head
	}
	for layer >= 0 {
		for i := 0; i < min(layerCount, len(now.level)); i++ {
			lastNode[i] = now
		}
		nex := now.level[layer]
		if nex == nil {
			layer--
		} else if nex.key > key {
			layer --
		} else if nex.key == key {
			break
		} else {
			now = nex
		}
	}
	//从上往下更新
	for i := layerCount - 1; i >= 0; i-- {
		temp := lastNode[i].level[i]
		lastNode[i].level[i] = &no
		no.level[i] = temp
	}
}
func (sl *SkipList) Tos() string {
	a := make([][]string, sl.size)
	for i := 0; i < len(a); i++ {
		a[i] = make([]string, len(sl.head.level))
	}
	ii := 0
	for i := &sl.head; i != nil; i = i.level[0] {
		for j := 0; j < len(i.level); j++ {
			a[ii][len(sl.head.level)-1-j] = fmt.Sprint(i.value)
		}
		ii++
	}
	s := make([]string, len(sl.head.level))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			s[j] += fmt.Sprintf("%5s", a[i][j])
		}
	}
	return strings.Join(s, "\n") + "\n============"
}
func TestOne(t *testing.T) {
	sl := NewSkipList()
	fmt.Println(sl.Tos())
	fmt.Println(sl.Get(0))
	sl.Put(0, "0")
	fmt.Println(sl.Tos())
	sl.Put(40, "40")
	fmt.Println(sl.Tos())
	sl.Put(28, "28")
	fmt.Println(sl.Tos())
	sl.Put(3, "3")
	fmt.Println(sl.Tos())
	sl.Put(7, "7")
	fmt.Println(sl.Tos())
	sl.Del(3)
	fmt.Println(sl.Tos())
	sl.Del(7)
	sl.Del(28)
	fmt.Println(sl.Tos())
	sl.Del(0)
	sl.Del(40)
	fmt.Println(sl.Tos())
}
func TestSecond(t *testing.T) {
	sl := NewSkipList()
	numberCount := 20
	for i := 0; i < numberCount; i++ {
		x := rand.Intn(100)
		fmt.Println("insert ", x)
		sl.Put(Key(x), Value(strconv.Itoa(x)))
		fmt.Println(sl.Tos())
	}
}
func TestThree(t *testing.T) {
	sl := NewSkipList()
	sl.Put(1, "1")
	sl.Del(1)
	fmt.Println(sl.Tos())
}
func main() {
	//TestOne(nil)
	TestThree(nil)
}
