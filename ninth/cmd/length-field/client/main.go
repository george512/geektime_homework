package main

import (
	"fmt"
	"geektime_homework/ninth/cmd/length-field/pkg"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("Hello, Hello, geroge?, %d", i)
		data, err := pkg.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
