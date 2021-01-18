package main

import "fmt"

func main()  {
	switch 1 + 1 {
	case 1:
		fmt.Println("1 + 1 = 1")
	case 2:
		fmt.Println("1 + 1 = 2")
	case 3:
		fmt.Println("1 + 1 = 3")
	default:
		fmt.Println("1 + 1 != 1, 1 + 1 != 2, 1 + 1 != 3")
	}

	//fallthrough
	fmt.Println("fallthrough example")
	switch{
	case true:
		fmt.Println("case 1 is true")
		fallthrough
	case false:
		fmt.Println("case 2 is false")
		fallthrough
	default:
		fmt.Println("default case")
	}
}