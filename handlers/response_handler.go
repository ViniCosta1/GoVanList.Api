package handlers

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vinicosta1/GoVanList.Api/config"
	"github.com/vinicosta1/GoVanList.Api/utils"
)

func GetResponse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := config.LoadConfig()
		jsonContent := utils.ReadJson(cfg.DataDir, cfg.FileName)
		fmt.Printf("JSON Content: %v", jsonContent)
		ctx.JSON(http.StatusOK, jsonContent)
	}
}	