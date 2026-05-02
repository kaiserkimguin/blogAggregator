package config


func (cfg *ConfigJson) SetUser(username string) error {
	// mutate the current config stuct
	cfg.CurrentUserName = username
	// write the struct to the config file
	return write(*cfg)
}
