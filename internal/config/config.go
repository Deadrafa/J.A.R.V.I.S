package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Token string `yaml:"token"`
}

func NewConfig(filename string) (*Config, error) {
	instance := &Config{}
	err := cleanenv.ReadConfig(filename, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
