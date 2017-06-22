package anime

import (
	"github.com/ElegantMonkey/neet/util"
)

// Search finds animes matching searchTerm on the list,
// using a non-strict search.
func Search(animes []Anime, searchTerm string) (out []Anime, err error) {
	re, err := util.MakeSearchRegex(searchTerm)
	for _, anime := range animes {
		if re.MatchString(anime.Name) {
			out = append(out, anime)
		}
	}

	return
}
