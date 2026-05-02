package config

import(
	"encoding/json"
	"os"
)

func write(cfg ConfigJson) error {
	// get config path
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}
	// marshal the data before writing them to the file
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0600)
}
