package middleware

import (
	"geektime_homework/fifth/pkg/hystrix"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Warpper(size, rate int, failedRate float64, duration time.Duration) gin.HandlerFunc {
	r := hystrix.NewRollWindow(size, rate, failedRate, duration)
	r.Launch()
	r.Monitor()
	r.ShowStatus()
	r.ShowTotalStatus()
	return func(c *gin.Context) {
		if r.FusingStatus() {
			c.String(http.StatusInternalServerError, "rejected by hystrix")
			c.Abort()
			return
		}
		c.Next()
		if c.Writer.Status() != http.StatusOK {
			r.RecordRes(false)
		} else {
			r.RecordRes(true)
		}
	}
}
