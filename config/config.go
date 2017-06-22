package config

import (
	"fmt"

	"github.com/shibukawa/configdir"
	yaml "gopkg.in/yaml.v2"
)

const (
	APPLICATION_NAME = "neet"
	CONFIG_FILE      = "config.yml"
)

type Config struct {
	AnimeFolder string `yaml:"anime_folder"`
	Player      string `yaml:"player"`
}

var Conf *Config

var defaultConfig = &Config{
	AnimeFolder: "/home/user/Videos",
	Player:      "mpv",
}

func Load() error {
	configDirs := configdir.New("", APPLICATION_NAME)
	folder := configDirs.QueryFolderContainsFile(CONFIG_FILE)
	if folder == nil {
		fmt.Println("WARNING: Config file doesn't exist, using (useless) default config.")
		Conf = defaultConfig
		return nil
	}

	data, err := folder.ReadFile(CONFIG_FILE)
	if err != nil {
		return err
	}

	err = Parse(data)
	if err != nil {
		return err
	}

	return nil
}

func Parse(data []byte) error {
	err := yaml.Unmarshal(data, &Conf)
	if err != nil {
		return err
	}

	return nil
}
