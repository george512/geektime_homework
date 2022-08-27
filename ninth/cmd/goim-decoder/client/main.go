package main

import (
	"fmt"
	"geektime_homework/ninth/pkg"
	"log"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		data := pkg.Encode(pkg.NewPack(uint16(i), uint32(i), uint32(i), []byte("hello"+strconv.Itoa(i))))
		if err != nil {
			log.Print("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
