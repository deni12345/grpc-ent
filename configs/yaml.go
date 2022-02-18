package configs

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	NewDecoder = yaml.NewDecoder
)

type yamlConfigure struct {
	path string
}

func NewYamlConfig(filepath string) *yamlConfigure {
	return &yamlConfigure{
		path: filepath,
	}
}
func (yc *yamlConfigure) Load(_ context.Context, value *values) error {
	file, err := os.Open(filepath.Clean(yc.path))
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error on closing file: %s", err)
		}
	}()
	if err := NewDecoder(file).Decode(value); err != nil {
		log.Printf("error on decode yaml file: %s", err)
		return err
	}
	return nil
}
