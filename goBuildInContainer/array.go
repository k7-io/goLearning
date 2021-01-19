package main

import "fmt"

func main()  {
	//array declaration
	var student1 [3]string
	fmt.Println(student1)
	fmt.Println(len(student1))
	fmt.Println(cap(student1))

	//array initialization and traversal
	var student2 = [...]string{"Tom", "Ben", "Peter", "Alice"}
	fmt.Println(student2)
	for k, v := range student2{
		fmt.Println("array index:", k, "value:", v)
	}

	//range
	var num = [...]int{1, 2, 3, 4}
	for k, v := range num{
		fmt.Println("variable k:", k, "variable v:", v)
	}
}