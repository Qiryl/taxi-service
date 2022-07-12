package config

import "github.com/kelseyhightower/envconfig"

type GrpcConfig struct {
	Port string
}

func GetEnvGrpcConfig() (*GrpcConfig, error) {
	c := new(GrpcConfig)
	err := envconfig.Process("GRPC", c)
	return c, err
}
