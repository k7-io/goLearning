package main

import "fmt"

func checkType(t interface{}, ok bool)  {
	if ok {
		fmt.Println("assert success")
	}else{
		fmt.Println("assert failed")
	}
	fmt.Printf("variable t type = %T, value = %v\n", t, t)
}

type Person interface {
	Speak()
}

type Student1 struct {
	Name string
}

func (st *Student1) Speak(){
	fmt.Println("My name is ", st.Name)
}

type Person1 struct {
	Name string
	Age int
}

type IPhoner interface {
	Call() error
	Video() error
	Game() error
}

type Apple interface {
	Call() error
	Video() error
}

type Huawei interface {
	Call() error
	Game() error
}

type Phone struct {
	Name string
}

func (p *Phone) Call() error {
	fmt.Println(p.Name, "Start Call")
	return nil
}

func (p *Phone) Video() error{
	fmt.Println(p.Name, "Start Video")
	return nil
}

func (p *Phone) Game() error{
	fmt.Println(p.Name, "Start Game")
	return nil
}

func main()  {
	var X interface{} = 1

	fmt.Println("first assert:")
	t0, ok := X.(string)
	checkType(t0, ok)

	fmt.Println("second assert")
	t1, ok := X.(float64)
	checkType(t1, ok)

	fmt.Println("third assert")
	t2, ok := X.(int)
	checkType(t2, ok)

	var student interface{} = Student1{"Xiao Ming"}
	fmt.Println("first assert")
	t3, ok := student.(string)
	checkType(t3, ok)
	fmt.Println("second assert")
	t4, ok := student.(Person)
	checkType(t4, ok)

	s := make([]interface{}, 3)
	s[0] = 1
	s[1] = "str"
	s[2] = Person1{"zhang san", 20}
	for index, data := range s{
		if value, ok := data.(int); ok == true {
			fmt.Printf("s[%d] Type = int, Value = %d\n", index, value)
		}
		if value, ok := data.(string); ok == true {
			fmt.Printf("s[%d] Type = string, Value = %s\n", index, value)
		}
		if value, ok := data.(Person1); ok == true {
			fmt.Printf("s[%d] Type = Person1, Person1.Name = %v, Person1.Age = %d\n", index, value.Name, value.Age)
		}
	}

	for index, data := range s{
		switch value :=  data.(type) {
		case int:
			fmt.Printf("s[%d] Type = int, Value = %d\n", index, value)
		case string:
			fmt.Printf("s[%d] Type = string, Value = %s\n", index, value)
		case Person1:
			fmt.Printf("s[%d] Type = Person1, Person1.Name = %v, Person1.Age = %d\n", index, value.Name, value.Age)
		}
	}

	var apple Apple = &Phone{"apple"}
	var huawei Huawei = new(Phone)
	apple.Video()
	huawei.Game()
}