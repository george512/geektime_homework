package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"geektime_homework/ninth/pkg"
	"io"
	"log"
	"net"
)

// socket_stick/server2/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// 获取包体长度
		lengthByte, err := reader.Peek(pkg.PackageLengthSize())
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Print("Peek failed:", err)
			return
		}

		buffer := bytes.NewBuffer(lengthByte)
		var size uint32
		if err := binary.Read(buffer, binary.BigEndian, &size); err != nil {
			log.Print("Read failed", err)
			return
		}
		if uint32(reader.Buffered()) < size {
			log.Print("error Size")
			return
		}
		data := make([]byte, size)
		if _, err := reader.Read(data); err != nil {
			log.Print("Read failed", err)
			return
		}
		pack, err := pkg.Decode(data)
		if err != nil {
			log.Print("Decode failed", err)
			return
		}
		log.Print("receive: ", string(pack.Body))
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
