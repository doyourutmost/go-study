package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
	}
	r.Use(func(context *gin.Context) {
		s := time.Now()
		context.Next()
		// log  path、status、latency
		logger.Info("incoming request",
			zap.String("string", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(context *gin.Context) {
		context.Set(keyRequestId, rand.Int())
		context.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if requestId, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = requestId
		}
		c.JSON(http.StatusOK, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
