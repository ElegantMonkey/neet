package main

import (
	"github.com/ElegantMonkey/neet/config"
	"github.com/ElegantMonkey/neet/commands"

	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	// Set to false to enable some messages.
	if true {
		log.SetOutput(ioutil.Discard)
	}

	err := config.Parse("config/default.yml")
	if err != nil {
		fmt.Printf("Error while parsing config: %v\n", err)
		return
	}

	commands.Parse()
}