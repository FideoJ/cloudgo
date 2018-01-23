package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
)

func main() {
	// get the port
	port := getPort()
	// create a new server
	server := createServer()
	// listen and serve on 0.0.0.0:${port}
	server.Run(":" + port)
}

func getPort() string {
	port := "8080"

	// get port from flag
	pPort := flag.StringP("port", "p", "8080", "PORT for cloudgo listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	return port
}

func createServer() *gin.Engine {
	// Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)
	// create a default server
	r := gin.Default()
	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		// reply with JSON
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   "hi " + name + "!",
		})
	})
	return r
}
