package main

import "fmt"

func main()  {
	a := 1
	b := 2
	fmt.Println("a + b:", a + b)
	fmt.Println("a - b:", a - b)
	fmt.Println("a * b:", a * b)
	fmt.Println("a / b:", a / b)
	fmt.Println("a % b:", a % b)
	a++
	fmt.Println("a++:", a)
	b--
	fmt.Println("b--:", b)
}