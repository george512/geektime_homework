package main

import (
	"geektime_homework/fifth/pkg/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	size       int           = 10
	rate       int           = 50
	failedRate float64       = 0.6
	duration   time.Duration = time.Second * 5
)

func main() {
	c := gin.Default()
	c.GET("/api/up/v1", middleware.Warpper(size, rate, failedRate, duration), upHandler)
	c.Run(":9000")
}

func upHandler(c *gin.Context) {
	res, err := http.Get("http://localhost:8000/api/down/v1")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, string(data))
		return
	}

	c.String(res.StatusCode, string(data))
}
