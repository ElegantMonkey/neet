package commands

import (
	"github.com/ElegantMonkey/neet/anime"
	"gopkg.in/alecthomas/kingpin.v2"

	"fmt"
)

var (
	add = kingpin.Command("add", "Add an anime")
	addName = add.Arg("name", "anime name").Required().String()
	addFolder = add.Arg("folder", "anime folder").Required().ExistingDir()
	addCurrentEpisode = add.Arg("current_episode", "current episode").Int()
)

func Add() {
	currentEpisode := 1
	if *addCurrentEpisode > 0 {
		currentEpisode = *addCurrentEpisode
	}

	a := anime.Anime{
		Folder: *addFolder,
		Name: *addName,
		CurrentEpisode: currentEpisode,
	}
	err := a.SaveConfig()
	if err == nil {
		fmt.Println("Anime added successfully!")
	} else {
		fmt.Printf("Failed to create anime config file: %v\n", err)
	}
}
