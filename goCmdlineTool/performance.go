package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

const Num int = 10000

func main() {
	f, _ := os.Create("cpu_file.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var str string
	for i := 0; i < Num; i++ {
		fmt.Sprintf("%s%d", str, i)
	}
}
