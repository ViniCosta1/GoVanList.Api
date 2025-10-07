package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/vinicosta1/GoVanList.Api/models"
)

func ReadJson(path, fileName string) models.Result {
	file := path + fileName

	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Erro na leitura do JSON: %v", err)
	}

	var r models.Result
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&r)

	return r
}