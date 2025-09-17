package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println(">>> init gin success...")

	// Define a simple hi gin endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": ">>> Hi Gin...",
		})
	})

	// Start the server on port 9000
	fmt.Println(">>> Start server...")
	r.Run(":9000")
}
