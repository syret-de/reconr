package internal

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	WorkPath  string   `yaml:"workPath"`
	MountPath string   `yaml:"mountPath"`
	Logfile   string   `yaml:"logfile"`
	Target    string   `yaml:"target"`
	Scope     []string `yaml:"scope"`
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

func (c *Config) getWorkPath() string {
	return c.WorkPath
}

func (c *Config) getMountPath() string {
	return c.MountPath
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
