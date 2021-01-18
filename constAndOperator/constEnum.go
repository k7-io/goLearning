package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c = "hello world"
		d
		e = iota
	)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
