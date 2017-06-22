package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"fmt"
	"strconv"
)

var (
	set      = kingpin.Command("set", "change anime settings")
	setName  = set.Arg("anime_name", "anime name").Required().String()
	setField = set.Arg("field", "field to be set (name, episode)").Required().Enum("name", "episode")
	setValue = set.Arg("value", "value to be set").Required().String()
)

func Set() {
	anime, err := SelectAnime(*setName)
	if err != nil {
		return
	}

	switch *setField {
	case "name":
		anime.Name = *setValue
	case "episode":
		ep, err := strconv.Atoi(*setValue)
		if err != nil {
			fmt.Printf("Invalid number: %v.", err)
			return
		}

		anime.CurrentEpisode = ep
	}
	anime.SaveConfig()
}
