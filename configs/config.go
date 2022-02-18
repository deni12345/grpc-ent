package configs

import (
	"context"
	"fmt"
	"log"
)

var (
	Values values
)

const local = "application"

type values struct {
	Host     string `yaml:"host" envconfig:"DB_ENDPOINT"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int64  `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func LoadConfigs() {
	Values = loadConfigurationValues(fmt.Sprintf("configs/%s.yaml", local))
}

func loadConfigurationValues(filepath string) values {
	v := values{}
	var ctx = context.Background()
	err := NewYamlConfig(filepath).Load(ctx, &v)
	if err != nil {
		log.Panicf("failed to load configuration: %w", err)
	}
	return v
}
