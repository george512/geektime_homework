package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for j := 0; j < 5; j++ {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				Client()
				wg.Done()
			}()
		}
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func Client() {
	res, err := http.Get("http://localhost:9000/api/up/v1")
	if err != nil {
		log.Print(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("%d: %s", res.StatusCode, string(data))
}
