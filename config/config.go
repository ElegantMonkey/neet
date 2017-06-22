package config

import (
	"fmt"

	"github.com/shibukawa/configdir"
	yaml "gopkg.in/yaml.v2"
)

// Constants used for the config paths.
const (
	ApplicationName = "neet"
	ConfigFile      = "config.yml"
)

// Config contains common configurable fields, used across the application.
type Config struct {
	AnimeFolder string `yaml:"anime_folder"`
	Player      string `yaml:"player"`
}

// Conf is the main Config instance.
var Conf *Config

var defaultConfig = &Config{
	AnimeFolder: "/home/user/Videos",
	Player:      "mpv",
}

// Load tries to read a config from the common system places.
// If no config file is found, loads a (useless) default config.
// Errors if there was a problem with the loaded config file.
func Load() error {
	configDirs := configdir.New("", ApplicationName)
	folder := configDirs.QueryFolderContainsFile(ConfigFile)
	if folder == nil {
		fmt.Println("WARNING: Config file doesn't exist, using (useless) default config.")
		Conf = defaultConfig
		return nil
	}

	data, err := folder.ReadFile(ConfigFile)
	if err != nil {
		return err
	}

	err = Parse(data)
	if err != nil {
		return err
	}

	return nil
}

// Parse a data array into a Config struct.
func Parse(data []byte) error {
	err := yaml.Unmarshal(data, &Conf)
	if err != nil {
		return err
	}

	return nil
}
