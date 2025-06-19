package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	TelegramToken      string `yaml:"telegrem_token"`
	GigaChatURL        string `yaml:"gigachat_url"`
	GigaChatToken      string `yaml:"gigachat_token"`
	Bearer             string `yaml:"bearer"`
	Model              string `yaml:"model"`
	Role               string `yaml:"role"`
	CalendarServiceURL string `yaml:"calendar_url"`
}

func NewConfig(filename string) (*Config, error) {
	instance := &Config{}
	err := cleanenv.ReadConfig(filename, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
