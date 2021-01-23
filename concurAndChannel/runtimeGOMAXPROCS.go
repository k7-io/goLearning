package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	n := runtime.GOMAXPROCS(1)
	fmt.Println("previos cpu is:", n)

	last := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {
			a := 999999 ^ 9999999
			a = a + 1
		}()
	}
	now := time.Now()
	fmt.Println(now.Sub(last))
}


