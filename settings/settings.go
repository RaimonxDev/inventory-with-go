package settings

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"HOST"`
	Port     int    `yaml:"PORT"`
	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	Name     string `yaml:"NAME"`
}

type Settings struct {
	Port string         `yaml:"PORT"`
	DB   DatabaseConfig `yaml:"DATABASE"`
}

func New() (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil

}
