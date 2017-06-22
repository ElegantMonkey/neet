package anime

import (
	"github.com/ElegantMonkey/neet/config"
	"path/filepath"
)

func ScanFolder(path string) (out []Anime, err error) {
	// Find config files
	matches, err := filepath.Glob(filepath.Join(path, "*", "neet.yml"))
	if err != nil {
		return
	}

	for _, match := range matches {
		anime := Anime{
			Folder: filepath.Dir(match),
		}
		anime.LoadConfig(match)

		out = append(out, anime)
	}

	return
}

func ScanFolderFromConfig() (out []Anime, err error) {
	out, err = ScanFolder(config.Conf.AnimeFolder)
	return
}
