package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"app"`
	HTTP struct {
		Port                int64 `yaml:"port"`
		MaxMultipartSizeMib int64 `yaml:"max_multipart_size_mib"`
	} `yaml:"http"`
	Postgres struct {
		PoolMax int64  `yaml:"pool_max"`
		Dsn     string `yaml:"dsn"`
	} `yaml:"postgres"`
	Logger struct {
		Level string `yaml:"level"`
	}
}

func MustParse(path string) config {
	cfg, err := New(path)
	if err != nil {
		panic(err)
	}

	return cfg
}

func New(path string) (config, error) {
	cfg := config{}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return config{}, errors.Wrap(err, "cleanenv.ReadConfig")
	}

	return cfg, nil
}
