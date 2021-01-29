package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}
	//NopCloser to create a body which implement io.Closer
	body := ioutil.NopCloser(strings.NewReader("user=admin&pass=admin"))
	reg, err := http.NewRequest("POST", "http://www.baidu.com", body)
	if err != nil {
		fmt.Println(err)
	}
	regBody, err1 := ioutil.ReadAll(reg.Body)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(regBody))
	//POST need to add Header
	reg.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err2 := client.Do(reg)
	if err2 != nil {
		fmt.Println(err2)
	}
	//panic, Body closed
	regBody3, err3 := ioutil.ReadAll(resp.Request.Body)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(regBody3))

	//http.Post
	resp4, err4 := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencoded", strings.NewReader("user=admin&pass=admin"))
	if err4 != nil {
		fmt.Println(err4)
	}
	defer resp4.Body.Close()
	body5, err5 := ioutil.ReadAll(resp4.Body)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(string(body5))

	//http.PostForm
	resp6, err6 := http.PostForm("http://www.baidu.com", url.Values{"user": {"admin"}, "pass": {"admin"}})
	if err6 != nil {
		fmt.Println(err6)
	}
	defer resp6.Body.Close()
	body7, err7 := ioutil.ReadAll(resp6.Body)
	if err7 != nil {
		fmt.Println(err7)
	}
	fmt.Println(string(body7))
}
