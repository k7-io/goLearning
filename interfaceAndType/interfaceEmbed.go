package main

import "fmt"

type IPerson interface {
	Speak()
}

type IStudent interface {
	IPerson
	Study()
}

type ITeacher interface {
	IPerson
	Teach()
}

type Student struct {
	Name string
}

func (student *Student) Speak()  {
	fmt.Println("My name is:", student.Name)
}

func (student *Student) Study()  {
	fmt.Printf("%s is studying now\n", student.Name)
}

type Teacher struct {
	Name string
}

func (teacher *Teacher) Speak()  {
	fmt.Println("My name is:", teacher.Name)
}

func (teacher *Teacher) Teach()  {
	fmt.Println(teacher.Name, "is teaching now")
}

func main()  {
	var stu Student = Student{"Tom"}
	stu.Speak()
	stu.Study()
	var teacher *Teacher= &Teacher{"Mr liu"}
	teacher.Speak()
	teacher.Teach()
}