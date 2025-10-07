package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicosta1/GoVanList.Api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.GetResponse())
	return r
}