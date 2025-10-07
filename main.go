package main

import (
	"github.com/vinicosta1/GoVanList.Api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/response", handler.GetResponse)
	router.Run()
}