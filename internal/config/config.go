package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type ConfigFloodControl struct {
	TimeLimit        time.Duration `yaml:"TimeLimit"`
	MaxQuantityQuery int           `yaml:"MaxQuantityQuery"`
}

func New(path string) (*ConfigFloodControl, error) {
	t, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out ConfigFloodControl
	err = yaml.Unmarshal(t, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
