package main

import "fmt"

func main()  {
	fmt.Println("hello")
	goto sign //Goto is generally not recommended in code
	fmt.Println("invalid code")
	sign:
		fmt.Println("world")
}