package main

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.Run(":9000")
}

func indexHandler(c *gin.Context) {
	var osinfo = map[string]string{
		"version": runtime.Version(),
		"arch":    runtime.GOARCH,
		"os":      runtime.GOOS,
	}

	c.JSON(http.StatusOK, osinfo)
}
