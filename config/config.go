package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	AnimeFolder string `yaml:"anime_folder"`
	Player string `yaml:"player"`
}

var Conf *Config

func Parse(configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		return err
	}

	return nil
}
