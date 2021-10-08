package main

import (
	"os"
	"runtime/pprof"
)

func four() {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 100; j++ {
			Five()
		}
	}
}
func Five() {
	for i := 0; i < 100; i++ {

	}
}
func Four() {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 100; j++ {
			Five()
		}
	}
}
func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic("create file error")
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic("start cpu profile error")
	}
	four()
	Four()
	defer pprof.StopCPUProfile()
}
