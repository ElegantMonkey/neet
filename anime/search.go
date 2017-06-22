package anime

import (
	"github.com/ElegantMonkey/neet/util"
)

func Search(animes []Anime, searchTerm string) (out []Anime, err error) {
	re, err := util.MakeSearchRegex(searchTerm)
	for _, anime := range animes {
		if re.MatchString(anime.Name) {
			out = append(out, anime)
		}
	}

	return
}
