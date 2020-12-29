package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func ReturnData() (int, int){
	return 10, 20
}

func main(){
	var name string
	name, age := "Tom", 18
	fmt.Println("name:", name)
	fmt.Println("age:", age)

	a := 2.0
	b := 3.0
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a/b)
	fmt.Println("var a's type is:", reflect.TypeOf(a))
	fmt.Println("var b's type is:", reflect.TypeOf(b))

	c, _ := ReturnData()
	_, d := ReturnData()
	fmt.Println(c, d)

	str :=`This is first line
This is second line
This is third line`
	english := 'a'
	fmt.Println(str)
	fmt.Println(english)

	var p *int
	p = new(int)
	fmt.Println("the pointer p value is:", *p)
	*p = 10
	fmt.Println("the pointer p value is:", *p)

	fmt.Println("hello world!")

	strA := "012345"
	strB := "6789"
	var strC bytes.Buffer
	strC.WriteString(strA)
	strC.WriteString(strB)
	fmt.Println(strC.String())
	fmt.Println(reflect.TypeOf(strC))

	strT := "Go language haha"
	index := strings.Index(strT, "lang")
	fmt.Println(index)
	fmt.Println(strT[index:])

	bytes := []byte(strT)
	for i := 0; i < 2; i++{
		bytes[i] = ' '
	}
	fmt.Println(string(bytes))
}

