package main

import (
	"fmt"
	"regexp"
)

func main() {
	targetString := "hello world"
	targetString1 := "Hello world"
	matchString := "hello"
	matchString1 := "(?i)hello" // (?i) Case insensitive

	matchString2 := `\w{3}`
	targetString2 := "123"
	matchString3 := "an?"
	targetString3 := "n"
	//regexp.MatchString
	match, err := regexp.MatchString(matchString, targetString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(match) //true
	match1, err1 := regexp.MatchString(matchString, targetString1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(match1) //false
	match2, err2 := regexp.MatchString(matchString1, targetString1)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(match2) //true
	match3, err3 := regexp.MatchString(matchString2, targetString2)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(match3) //true
	match4, err4 := regexp.MatchString(matchString3, targetString3)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(match4) //false

	//FindStringIndex
	re := regexp.MustCompile(`(\w)+`)
	res := re.FindStringIndex(targetString)
	fmt.Println(res) //[0 5]
	fmt.Println(res[0])
	fmt.Println(res[1])

	//ReplaceAllString
	re1 := regexp.MustCompile(`o`)
	res1 := re1.ReplaceAllString(targetString, `O`)
	fmt.Println(res1)
}
