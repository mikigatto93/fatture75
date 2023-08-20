package controller

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	AccessToken   string `json:"access_token"`
	CompanyId     int    `json:"company_id"`
	ModelFilePath string `json:"model_file_path"`
}

func NewConfiguration(filePath string) Configuration {

	conf := Configuration{}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return conf //empty configuration
	}

	json.Unmarshal(data, &conf)
	return conf

}
