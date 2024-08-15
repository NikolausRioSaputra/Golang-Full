package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/consume/rest", consumeREST)
	r.GET("/consume/soap", consumeSOAP)
	r.Run(":8080")
}
