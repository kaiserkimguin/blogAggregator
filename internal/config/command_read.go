package config

import (
	"os"
	"encoding/json"
)

func Read() (ConfigJson, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return ConfigJson{}, nil 
	}
	// create empty return variable
	var conJ ConfigJson
	// try to read the config path
	rawFile, err := os.ReadFile(configPath)
	if err != nil {
		return ConfigJson{}, err
	}
	// marshal the data to json format
	if err = json.Unmarshal(rawFile, &conJ); err != nil {
		return ConfigJson{}, err
	}
	// return the json struct
	return conJ, nil
} 

