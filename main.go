package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Run(":8080") // localhost:8080
}
