package anime

import (
	"github.com/ElegantMonkey/neet/config"
	"path/filepath"
)

// ScanFolder searches through a folder, finding neet.yml configuration
// files and loading them into Anime structs. Returns a slice of Anime.
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

// ScanFolderFromConfig runs ScanFolder on the default
// "anime_folder" on the Config.
func ScanFolderFromConfig() (out []Anime, err error) {
	out, err = ScanFolder(config.Conf.AnimeFolder)
	return
}
