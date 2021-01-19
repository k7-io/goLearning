package main

import "fmt"

func passByValue(numPara int)  {
	fmt.Printf("numPara's address of passByValue:%p\n", &numPara)
	numPara = 100
}

func passByReference(numPara *int)  {
	fmt.Printf("numPara's address of passByReference:%p\n", &numPara)
	fmt.Printf("numPara point to address of passByReference:%p\n", numPara)
	*numPara = 100
}

func passByReferenceMap(mapNumReference map[int]int)  {
	fmt.Printf("mapNumReference's address of passByReferenceMap:%p\n", &mapNumReference)
	fmt.Printf("mapNumReference point to address of passByReferenceMap:%p\n", mapNumReference)
	mapNumReference[1] = 100
}

func main()  {
	num := 1
	fmt.Printf("num's address of main:%p\n", &num)
	passByValue(num)
	fmt.Println("num:", num)
	passByReference(&num)
	fmt.Println("num:", num)

	mapNum := map[int]int{1:10}
	fmt.Printf("mapNum's address of main:%p\n", &mapNum)
	fmt.Printf("mapNum point to address of main:%p\n", mapNum)
	fmt.Println("mapNum:", mapNum)
	passByReferenceMap(mapNum)
	fmt.Println("mapNum:", mapNum)

	//GO inner function in builtin.go
	//append 4, 5, 6 to sourceSlice
	sourceSlice := []int{1, 2, 3}
	sourceSlice = append(sourceSlice, 4, 5, 6)
	fmt.Println("sourceSlice:", sourceSlice)

	//copy elements in sourceSlice to targetSlice
	targetSlice := make([]int, 3)
	number := copy(targetSlice, sourceSlice)
	fmt.Println("number:", number)
	fmt.Println("targetSlice:", targetSlice)
	fmt.Println("targetSlice's length:", len(targetSlice), ",targetSlice's cap:", cap(targetSlice))

	//complex function
	c1 := complex(1, 2)
	fmt.Println("c1's real:", real(c1), ",c1's imag:", imag(c1))
}