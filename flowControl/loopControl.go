package main

import "fmt"

func main()  {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}

	//break
	fmt.Println("break example")
	i := 1
	BreakOuterLoop:
		for {
			for {
				if i > 5 {
					fmt.Println("jump outer for loop")
					break BreakOuterLoop
				}
				fmt.Println(i)
				i++
			}
		}

	//continue
	fmt.Println("continue example")
	ContinueOuterLoop:
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				if j == i {
					fmt.Println(i, j)
					continue ContinueOuterLoop
				}
			}
		}
}