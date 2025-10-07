package main

import (
	"github.com/vinicosta1/GoVanList.Api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}