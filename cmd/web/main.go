package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	version := "¯\\_(ツ)_/¯"
	if os.Getenv("VERSION") != "" {
		version = os.Getenv("VERSION")
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": version,
			"more":    "chello",
		})
	})

	port := "8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)

	r.Run(addr)
}
