package config

import "github.com/kelseyhightower/envconfig"

type PostgresConfig struct {
	Url    string
	Driver string
}

func GetEnvPostgresConfig() (*PostgresConfig, error) {
	c := new(PostgresConfig)
	err := envconfig.Process("DB", c)
	return c, err
}
