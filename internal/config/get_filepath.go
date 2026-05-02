package config

import(
	"path/filepath"
	"os"
)
func getConfigPath () (string, error) {
	// determining users home dir.
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// set entire path to config file
	configPath :=  filepath.Join(home, configFileName)
	// return the path
	return configPath, nil
}
