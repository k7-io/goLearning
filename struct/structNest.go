package main

import (
	"fmt"
	"encoding/json"
)

type BookGeneral struct {
	title, author string
	num, id int
}

type BookBorrow struct {
	BookGeneral
	borrowTime string
}

type BookNotBorrow struct {
	BookGeneral
	readTime string
}

type BookBorrowAnonymous struct {
	Book struct{
		title, author string
		num, id int
	}
	borrowTime string
}

type Result struct {
	Person []Person
}

type Person struct {
	Name, Age string
	Interests []Interests
}

type Interests struct {
	Interest []string
}

func main()  {
	bookBorrow := &BookBorrow{
		BookGeneral:BookGeneral{
			"GO language",
			"Tom",
			20,
			15268,
		},
		borrowTime: "30",
	}
	bookNotBorrow := new(BookNotBorrow)
	bookNotBorrow.title = "Python language"
	bookNotBorrow.author = "Peter"
	bookNotBorrow.num = 30
	bookNotBorrow.id = 16428
	bookNotBorrow.readTime = "60"
	fmt.Println(bookBorrow)
	fmt.Println(bookNotBorrow)

	//nest struct anonymous
	bookBorrowAnony := &BookBorrowAnonymous{
		Book: struct {
			title, author string
			num, id       int
		}{title: "C language", author: "Peter" , num: 20 , id: 13689},
		borrowTime: "36",
	}
	fmt.Println(bookBorrowAnony)

	//struct anonymous
	book1 := struct {
		title, author string
		num, id int
	}{"GO language", "Tom", 20, 15678}
	fmt.Println("title:", book1.title)
	fmt.Println("author:", book1.author)
	fmt.Println("num:", book1.num)
	fmt.Println("id:", book1.id)

	//struct anonymous use
	data := &struct {
		Code int
		Msg string
	}{}
	jsonData := `{"code": 200, "msg": ""}`
	if err := json.Unmarshal([]byte(jsonData), data); err != nil{
		fmt.Println(err)
	}
	fmt.Println("code:", data.Code)
	fmt.Println("msg:", data.Msg)
}
