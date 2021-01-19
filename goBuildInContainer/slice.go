package main

import "fmt"

func main()  {
	//Generate a new slice from the array
	var student = [...]string{"Tom", "Ben", "Peter"}
	var student1 = student[1:2]
	fmt.Println("array student:", student)
	fmt.Println("slice student1:", student1)
	fmt.Println("the second element's address of student:", &student[1])
	fmt.Println("the first element's address of student1:", &student1[0])
	fmt.Println("slice student1's length:", len(student1))
	fmt.Println("slice student1's cap:", cap(student1))

	//Generate a new slice from the slice
	var student2 = student[1:3]
	var student3 = student2[0:1]
	fmt.Println("array student", student)
	fmt.Println("slice student2:", student2)
	fmt.Println("slice student3:", student3)
	fmt.Println("the second element's address of student:", &student[1])
	fmt.Println("the first element's address of student2:", &student2[0])
	fmt.Println("the first element's address of student3:", &student3[0])
	fmt.Println("slice student2's length:", len(student2))
	fmt.Println("slice student2's cap:", cap(student2))
	fmt.Println("slice student3's length:", len(student3))
	fmt.Println("slice student3's cap:", cap(student3))

	//Generate a new slice directly
	var teacher = []string{"Tom", "Ben", "Peter"}
	fmt.Println("slice teacher:", teacher)
	fmt.Println("slice teacher's length:", len(teacher))
	fmt.Println("slice teacher's cap:", cap(teacher))
	fmt.Println("if teacher is empty:", teacher==nil)

	//make() initialize
	var nums []int
	nums = make([]int, 1, 1)
	fmt.Println("slice nums:", nums)
	fmt.Println("slice nums's length:", len(nums))
	fmt.Println("slice nums's cap:", cap(nums))
	fmt.Println("if nums is empty:", nums==nil)

	//append element for slice
	for i := 0; i < 8; i++ {
		nums = append(nums, i)
		fmt.Println("slice nums's length:", len(nums), "slice nums's cap:", cap(nums))
	}
	var teacher1 = teacher[0:1]
	fmt.Println("slice teacher:", teacher)
	fmt.Println("slice teacher1:", teacher1)
	teacher1 = append(teacher1, "Danny")
	fmt.Println("slice teacher1:", teacher1, "teacher1's length:", len(teacher1), "teacher1's cap:", cap(teacher1))
	fmt.Println("slice teacher:", teacher)

	//Traverse slice
	for k, v := range teacher{
		fmt.Println("slice index:", k, "value:", v)
	}

	//delete element for slice
	fmt.Println("slice teacher:", teacher)
	teacher = append(teacher[0:1], teacher[2:]...)
	fmt.Println("slice teacher:", teacher, "teacher's length:", len(teacher), "teacher's cap:", cap(teacher))

	//clear slice teacher
	teacher = teacher[0:0]
	fmt.Println("slice teacher:", teacher, "teacher's length:", len(teacher), "teacher's cap:", cap(teacher))
}