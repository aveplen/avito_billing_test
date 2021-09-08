package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server      `yaml:"server"`
		Postgres    `yaml:"postgres"`
		CurrencyApi `yaml:"currency_api"`
	}

	Server struct {
		BindAddr string `yaml:"bind_addr"`
	}

	Postgres struct {
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     string `yaml:"port" env:"DB_PORT"`
		User     string `yaml:"user" env:"DB_USERNAME"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		DBName   string `yaml:"dbname" env:"DB_NAME"`
		SSLMode  string `yaml:"sslmode" env:"DB_SSLMODE"`
	}

	CurrencyApi struct {
		Key string `yaml:"key"`
	}
)

func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
