package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err1 := client.Do(request)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(response.StatusCode)

	res, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(res))

	//Cookie
	cookie := &http.Cookie{Name: "userId", Value: strconv.Itoa(12345)}
	request.AddCookie(cookie)
	request.AddCookie(&http.Cookie{Name: "session", Value: "YWRtaW4="})
	response3, err3 := client.Do(request)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(response3.Request.Cookies())

	//http.Get  Quick launch Get request
	response4, err4 := http.Get("http://www.baidu.com")
	if err4 != nil {
		fmt.Println(err4)
	}
	defer response4.Body.Close()
	body, _ := ioutil.ReadAll(response4.Body)
	fmt.Println(string(body))

	//UA to Bypass anti climbing strategy
	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK, utf-8;q=0.7, *;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	request.Header.Set("Accept-Language", "zh-CN, zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) chrome/72.0.3626.121")
	response5, err5 := client.Do(request)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Printf("%#v", response5.Request.Header)
}
