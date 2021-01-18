package main

import (
	"fmt"
	"runtime"
)

func main()  {
	a := 101
	if a > 100{
		fmt.Println("a > 100, a = ", a)
	}else if a == 100{
		fmt.Println("a == 100")
	}else{
		fmt.Println("a < 100, a = ", a)
	}

	if num := runtime.NumCPU(); num >= 1{
		fmt.Println("The process's cpu num is ", num)
	}
}
