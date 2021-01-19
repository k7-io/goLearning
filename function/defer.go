package main

import (
	"fmt"
	"net"
)

func tcpSend()  {
	//use defer to release resource
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err == nil {
		defer conn.Close()
		fmt.Println("remote address:", conn.RemoteAddr())
	}else{
		fmt.Println("error:", err)
	}
}

func main()  {
	fmt.Println("Start now")
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	defer fmt.Println("third defer")
	fmt.Println("end")

	tcpSend()
}