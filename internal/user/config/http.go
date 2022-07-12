package config

import "github.com/kelseyhightower/envconfig"

type HttpConfig struct {
	Port string
}

func GetEnvHttpConfig() (*HttpConfig, error) {
	c := new(HttpConfig)
	err := envconfig.Process("HTTP", c)
	return c, err
}
