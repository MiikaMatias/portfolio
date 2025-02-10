package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Name            string `yaml:"name"`
		WorkingDir      string `yaml:"workingDir"`
		BackendPort     string `yaml:"backendPort"`
		EncodingUrl     string `yaml:"encodingUrl"`
		EncodingModel   string `yaml:"encodingModel"`
		GeneratingUrl   string `yaml:"generatingUrl"`
		GeneratingModel string `yaml:"generatingModel"`
	} `yaml:"app"`
}

func CastConfig(configFilePath string, configType *Config) error {
	file, err := os.Open(configFilePath)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(configType); err != nil {
		return err
	}

	return nil
}
