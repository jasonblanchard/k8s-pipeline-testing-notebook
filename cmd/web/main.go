package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "¯\\_(ツ)_/¯",
		})
	})

	port := "8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)

	r.Run(addr)
}
