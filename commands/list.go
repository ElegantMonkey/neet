package commands

import (
	"github.com/ElegantMonkey/neet/anime"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"

	"fmt"
	"os"
)

var (
	list = kingpin.Command("list", "list animes")
)

func List() {
	animes, err := anime.ScanFolderFromConfig()

	if err != nil {
		fmt.Printf("Failed to scan anime folder: %v\n", err)
		return
	}

	fmt.Printf("Found %d animes\n", len(animes))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Progress"})
	for _, anime := range animes {
		progress := fmt.Sprintf("%d/%d", anime.CurrentEpisode, anime.NumEpisodes())
		table.Append([]string{anime.Name, progress})
	}
	table.Render()
}