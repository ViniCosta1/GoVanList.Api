package model

const (
	DataDir  = "E:\\Projects\\GoVanList.Api\\data\\"
	JsonFile = DataDir + "response.json"
)

type ResponseData struct {
	Dia       string   `json:"day"`
	Pessoas   []string `json:"people"`
	Pontos    []string `json:"places"`
	Aviso     string   `json:"warning"`
	ApenasIda bool     `json:"just_go"`
}

type Result struct {
	Result ResponseData `json:"data"`
}
