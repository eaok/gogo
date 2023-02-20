package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()
	router.Use(TlsHandler())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello index")
	})

	router.RunTLS(":8080", "/etc/nginx/cert/server.crt", "/etc/nginx/cert/server.key")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
