package anime

import (
	"github.com/ElegantMonkey/neet/util"
	"github.com/ElegantMonkey/neet/util/title_parser"

	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

// Episodes stores info about an Episode. Currently only holds the
// episode number, but might contain more metadata in the future.
type Episode struct {
	Path string
	// Currently only an int, but this might change (because why follow standards right)
	ID int
}

// EpisodeFromPath builds a Episode struct from a path,
// reading some metadata.
func EpisodeFromPath(path string) Episode {
	return Episode{
		Path: path,
		ID:   title_parser.Episode(path),
	}
}

// Episodes goes through the anime folder, returning a list of episodes
// currently available.
func (a *Anime) Episodes() (out []Episode, err error) {
	files, err := ioutil.ReadDir(a.Folder)

	log.Printf("Found %d files for %s\n", len(files), a.Name)
	if err != nil {
		return
	}

	for _, file := range files {
		log.Printf("File name: %s\n", file.Name())
		if !file.IsDir() && util.IsVideoFile(file.Name()) {
			log.Println("making ep struct")
			episode := EpisodeFromPath(filepath.Join(a.Folder, file.Name()))
			log.Printf("%v\n", episode)
			out = append(out, episode)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})

	return
}

// NumEpisodes returns the highest-number episode found. Returns 0 on error or if no episodes were found.
func (a *Anime) NumEpisodes() int {
	eps, err := a.Episodes()
	if err != nil {
		log.Printf("Warn while getting ep num: %v\n", err)
		return 0
	}

	max := 0
	for _, ep := range eps {
		if ep.ID > max {
			max = ep.ID
		}
	}

	return max
}

func getEpisode(episodes []Episode, episode int) *Episode {
	// The *Episode return type is only there to allow us
	// to return nil - maybe we should use an error?
	for _, ep := range episodes {
		if ep.ID == episode {
			return &ep
		}
	}

	return nil
}

// GetEpisode gets a episode from a episode list
// Probably isn't that necessary, but whatever.
func (a *Anime) GetEpisode(episode int) *Episode {
	// FIXME Possibly a silent error?
	eps, err := a.Episodes()
	if err != nil {
		return nil
	}

	return getEpisode(eps, episode)
}
