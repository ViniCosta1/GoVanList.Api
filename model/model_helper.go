package model

import (
	"os"
	"log"
	"encoding/json"
	"bytes"
)

func ReadJson() Result {
	jsonFile, err := os.ReadFile(JsonFile)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}

	var r Result
	json.NewDecoder(bytes.NewBuffer(jsonFile)).Decode(&r)

	return r
}