package internal

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	MountWork     string   `yaml:"mountWork"`
	WorkPath      string   `yaml:"workPath"`
	Logfile       string   `yaml:"logfile"`
	ScopeFileName string   `yaml:"scopeFileName"`
	Target        string   `yaml:"target"`
	Scope         []string `yaml:"scope"`
	ConfigPath    string   `yaml:"configPath"`
	MountConfig   string   `yaml:"mountConfig"`
}

func NewConfig(path string) (Config, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) GetWorkPath() string {
	return c.WorkPath
}

func (c *Config) GetMountWork() string {
	return c.MountWork
}

func (c *Config) GetLogPath() string {
	return c.Logfile
}

func (c *Config) GetTarget() string {
	return c.Target
}

func (c *Config) GetScope() []string {
	return c.Scope
}

func (c *Config) GetConfigPath() string {
	return c.ConfigPath
}

func (c *Config) GetMountConfig() string {
	return c.MountConfig
}

func (c *Config) GetScopeFileName() string {
	return c.ScopeFileName
}
