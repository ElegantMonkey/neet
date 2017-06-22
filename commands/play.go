package commands

import (
	"github.com/ElegantMonkey/neet/anime"
	"github.com/ElegantMonkey/neet/util"
	"gopkg.in/alecthomas/kingpin.v2"

	"fmt"
	"log"
)

var (
	play = kingpin.Command("play", "Play an anime")
	playName = play.Arg("name", "anime name").Required().String()
	playMode = play.Arg("mode", "play mode (current, next, default is next)").Default("next").Enum("current", "next")
)

func playEpisode(a anime.Anime, targetEpisode int) error {
	total := a.NumEpisodes();
	if targetEpisode > total {
		return fmt.Errorf("Target episode (%d) is bigger than the number of episodes (%d).\n", targetEpisode, total)
	}

	episode := a.GetEpisode(targetEpisode)
	if episode == nil {
		return fmt.Errorf("Episode %d does not exist!\n", targetEpisode)
	}

	util.LaunchPlayer(episode.Path)
	return nil
}

func playNext(a anime.Anime, targetEpisode int) {
	err := playEpisode(a, targetEpisode)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	total := a.NumEpisodes()
	if targetEpisode + 1 <= total {
		log.Printf("Setting current episode to %d of %d", targetEpisode + 1, total)
		a.CurrentEpisode = targetEpisode + 1

		a.SaveConfig()
		playNext(a, targetEpisode + 1)
	}
}

func Play() {
	log.Println("Selecting anime")
	anime, err := SelectAnime(*playName)
	if err != nil {
		log.Println("Anime not found!")
		return
	}

	switch *playMode {
	case "current":
		err = playEpisode(anime, anime.CurrentEpisode)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return
		}
	case "next":
		playNext(anime, anime.CurrentEpisode + 1)
	default:
		// This shouldn't happen, because kingpin handles it
		// but just in case...
		panic("Invalid play mode.")
	}
}
