package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

const (
	_configName = "rit.yaml"
)

type Credential struct {
	Id       string `yaml:"id"`
	Secret   string `yaml:"secret"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Credential Credential `yaml:"credential"`
}

func (c *Config) merge(cfg *Config) {
	if len(c.Credential.Id) == 0 && len(cfg.Credential.Id) != 0 {
		c.Credential.Id = cfg.Credential.Id
	}

	if len(c.Credential.Secret) == 0 && len(cfg.Credential.Secret) != 0 {
		c.Credential.Secret = cfg.Credential.Secret
	}

	if len(c.Credential.Username) == 0 && len(cfg.Credential.Username) != 0 {
		c.Credential.Username = cfg.Credential.Username
	}

	if len(c.Credential.Password) == 0 && len(cfg.Credential.Password) != 0 {
		c.Credential.Password = cfg.Credential.Password
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func LoadConfig() *Config {
	var defaultCfg = DefaultConfig()
	configFilePath, err := xdg.SearchConfigFile(filepath.Join("rit", _configName))
	if err != nil {
		return defaultCfg
	}

	content, err := os.ReadFile(configFilePath)
	if err != nil {
		// TODO: add log
		return defaultCfg
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return defaultCfg
	}

	config.merge(DefaultConfig())
	return &config
}
