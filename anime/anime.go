package anime

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Anime struct {
	Folder string `yaml:"-"` // Hide the Folder field from yaml, as it's the same folder as the anime cfg file.
	Name string `yaml:"name,omitempty"`
	CurrentEpisode int `yaml:"current_episode,omitempty"`
}

func (a *Anime) LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	return nil
}

// SaveConfig saves the config to the default path (Folder/neet.yml)
func (a *Anime) SaveConfig() error {
	path := filepath.Join(a.Folder, "neet.yml")

	data, err := yaml.Marshal(&a)
	if err != nil {
		return err
	}

	// Default perms are u+rw,g+r,o+r
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

