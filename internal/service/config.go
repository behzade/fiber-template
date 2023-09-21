package service

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Service struct {
		Name string `yaml:"name" envconfig:"SERVICE_NAME"`
		Env  string `yaml:"env" envconfig:"SERVICE_ENV"`
	} `yaml:"service"`
	Server struct {
		Port int `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host" envconfig:"DB_HOST"`
		Port int    `yaml:"port" envconfig:"DB_PORT"`
		Name string `yaml:"name" envconfig:"DB_NAME"`
		User string `yaml:"user" envconfig:"DB_USER"`
		Pass string `yaml:"pass" envconfig:"DB_PASS"`
	} `yaml:"database"`
	OpenTelemtry struct {
		Host     string `yaml:"host" envconfig:"OPENTELEMTRY_HOST"`
		Insecure bool   `yaml:"insecure" envconfig:"OPENTELEMTRY_INSECURE"`
	} `yaml:"opentelemtry"`
}

func LoadConfig() (*Config, error) {
	cfg := Config{}

	err := readYamlConfig(&cfg)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readYamlConfig(cfg *Config) error {
	f, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}
