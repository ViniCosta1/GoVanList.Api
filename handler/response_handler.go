package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicosta1/GoVanList.Api/model"
)

func GetResponse(c *gin.Context)  {
	c.JSON(http.StatusOK, model.ReadJson())
}