package main

import (
	"fmt"

	"urlshortener/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/encode", handler.CreateShortUrl)
	r.POST("/decode", handler.DecodeShortUrl)

	err := r.Run(":7070")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
