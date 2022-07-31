package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

const (
	successRate float64 = 0.2
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	c := gin.Default()
	c.GET("/api/down/v1", downHandler)
	c.Run(":8000")
}

func downHandler(c *gin.Context) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	if !rejectOrNot() {
		c.String(http.StatusInternalServerError, "reject from downStream")
		return
	}
	c.String(http.StatusOK, "approve from downStream")
}

func rejectOrNot() bool {
	return rand.Float64() < successRate
}
