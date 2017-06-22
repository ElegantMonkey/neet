package commands

import (
	"errors"
	"fmt"
	"github.com/ElegantMonkey/neet/anime"
)

func SelectAnime(searchTerm string) (a anime.Anime, err error) {
	animes, err := anime.ScanFolderFromConfig()
	if err != nil {
		fmt.Printf("Error while loading animes from folder: %v\n")
		return
	}

	// Search for animes with this name, and display a
	// selection screen when possible
	results, err := anime.Search(animes, *setName)
	if err != nil {
		fmt.Printf("Error while searching for animes matching '%s': %v\n", *setName, err)
		return
	}

	a = results[0]
	if len(results) > 1 {
		// TODO implement choice
		fmt.Println("More than one anime matches the search, aborting.")
		err = errors.New("More than one anime matches.")
		return
	}

	return
}
