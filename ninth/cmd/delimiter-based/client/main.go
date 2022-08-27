package main

import (
	"fmt"
	"log"
	"net"
)

const (
	Delimeter = '\n'
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Print("Dial failed:", err)
		return
	}
	defer conn.Close()

	// sending message
	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("hello, world %d", i)
		_, err = conn.Write(patch(content))
		if err != nil {
			log.Printf("Write failed at %d, error:%s", i, err)
			return
		}
	}
}

func patch(content string) []byte {
	data := []byte(content)
	data = append(data, Delimeter)
	return data
}
