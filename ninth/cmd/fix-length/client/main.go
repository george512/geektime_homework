package main

import (
	"fmt"
	"log"
	"net"
)

const (
	BuffSize = 1024
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Print("Dial failed:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		content := fmt.Sprintf("hello, world, %d", i)
		_, err = conn.Write(patch(content))
		if err != nil {
			log.Printf("Write failed at %d, error:%s", i, err)
			return
		}
	}
}

// patching empty byte into origin message
func patch(message string) []byte {
	res := make([]byte, BuffSize)
	copy(res, []byte(message))
	return res
}
