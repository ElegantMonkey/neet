package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func Parse() {
	switch kingpin.Parse() {
	case "add":
		Add()
	case "list":
		List()
	case "set":
		Set()
	case "play":
		Play()
	}
}
