package main

import (
	"github.com/ElegantMonkey/neet/commands"
	"github.com/ElegantMonkey/neet/config"

	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Set to false to enable some messages.
	if true {
		log.SetOutput(ioutil.Discard)
	}

	err := config.Load()
	if err != nil {
		fmt.Printf("Error while parsing config: %v\n", err)
		return
	}

	commands.Parse()
}
