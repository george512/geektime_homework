package main

import (
	"bytes"
	"io"
	"log"
	"net"
)

const (
	Delimeter = '\n'
	BuffSize  = 1024
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

		// pointer for read buffer
		var start int
		var end int
		for k, v := range result.Bytes() {
			// when byte equals to defined delimeter, then set to end pointer
			if v == Delimeter {
				end = k
				log.Printf("recevie: %v", string(result.Bytes()[start:end]))
				// move start pointer
				start = end + 1
			}
		}
		result.Reset()
	}
}
