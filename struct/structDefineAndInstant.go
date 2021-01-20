package main

import "fmt"


type Book struct {
	title, author string
	num, id int
}

func main()  {
	//standard instance
	var book1 Book
	fmt.Println(book1)

	//new function instance
	book2 := new(Book)
	fmt.Println(book2)

	//& instance
	book3 := &Book{}
	fmt.Println(book3)

	//visit members
	book3.title = "GO language"
	book3.author = "Tom"
	book3.num = 20
	book3.id = 152368
	fmt.Println("title:", book3.title)
	fmt.Println("author:", book3.author)
	fmt.Println("num:", book3.num)
	fmt.Println("id:", book3.id)

	//struct map initialize
	book4 := Book{// &Book
		title: "C language",
		author: "Peter",
		num: 10,
		id: 152356,
	}
	fmt.Println("title:", book4.title)
	fmt.Println("author:", book4.author)
	fmt.Println("num:", book4.num)
	fmt.Println("id:", book4.id)

	//struct list initialize
	book5 := &Book{ //Book
		"Java",
		"Tom",
		30,
		123456,
	}
	fmt.Println("title:", book5.title)
	fmt.Println("author:", book5.author)
	fmt.Println("num:", book5.num)
	fmt.Println("id:", book5.id)
}