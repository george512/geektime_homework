package main

import (
	"bytes"
	"io"
	"log"
	"net"
)

const (
	BuffSize = 1024
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Print("Listen failed:", err)
		return
	}

	log.Print("listening at 127.0.0.1:8080")

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Print("Accept failed:", err)
			continue
		}
		go handler(conn)
	}
}

func handler(c net.Conn) {
	defer c.Close()
	buf := make([]byte, BuffSize)
	result := bytes.NewBuffer(nil)
	for {
		n, err := c.Read(buf)

		if err == io.EOF {
			return
		}

		if err != nil {
			log.Print("Read failed:", err)
			return
		}
		result.Write(buf[0:n])
		log.Printf("recevie size[%d]: %v", n, result.String())
		result.Reset()
	}
}
