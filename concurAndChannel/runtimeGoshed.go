package main

import (
	"fmt"
	"runtime"
)

//Goshed(), It does not suspend the current goroutine, so execution resumes automatically.

func main()  {
	go func() {
		for i := 0; i < 3; i++{
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++{
		runtime.Gosched()//Abandon CPU to let other goroutine run
		fmt.Println("Main")
	}
}