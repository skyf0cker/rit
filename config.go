package main

type Credential struct {
	Id       string
	Secret   string
	Username string
	Password string
}

type Config struct {
	Credential Credential `json:"credential"`
}

func DefaultConfig() *Config {
	return &Config{}
}

func LoadConfig() *Config {
	// TODO: load config from XDG file
	return DefaultConfig()
}
