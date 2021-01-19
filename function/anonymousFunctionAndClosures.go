package main

import "fmt"

func addOne(i int) func() int {
	return func() int{
		i++
		return i
	}
}

func main()  {
	//func(data string){
	//	fmt.Println("hello " + data)
	//}("world")

	f1 := func(data string){
		fmt.Println("hello " + data)
	}
	f1("world")

	//closures's memory
	num := 1
	fmt.Printf("%p\n", &num)
	func(){
		num++
		fmt.Println("num:", num)
		fmt.Printf("%p\n", &num)
	}()
	func(){
		num++
		fmt.Println("num:", num)
		fmt.Printf("%p\n", &num)
	}()

	a1 := addOne(0)
	fmt.Println(a1())
	fmt.Println(a1())
	a2 := addOne(10)
	fmt.Println(a2())
	fmt.Printf("closure a1's address:%p\n", &a1)
	fmt.Printf("closure a2's address:%p\n", &a2)
}