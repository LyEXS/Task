package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type CommitType struct {
	Label string `yaml:"label"`
	Emoji string `yaml:"emoji"`
}

type Config struct {
	Editor      string       `yaml:"editor"`
	CommitTypes []CommitType `yaml:"commit_types"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
