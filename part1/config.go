package part1

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	GrpcPort string      `yaml:"grpc_port"`
	Mq       MqConfig    `yaml:"mq"`
	Redis    RedisConfig `yaml:"redis"`
}

type MqConfig struct {
	Broker string `yaml:"broker"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

func ReadConfig(env string) (cfg Config, err error) {
	absPath, err := filepath.Abs(fmt.Sprintf("files/%s.yaml", env))
	if err != nil {
		return
	}
	file, err := os.Open(absPath)
	if err != nil {
		return
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	err = d.Decode(&cfg)

	return
}
