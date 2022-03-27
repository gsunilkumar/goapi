package main

import (
	"github.com/gin-gonic/gin"
	"service"
)


func main() {
    router := gin.Default()
    router.POST("/process", service.HandleRequest)
	router.Run(":7004")
}
