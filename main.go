package main

import (
	"time"

	"github.com/bigveezus/my-cli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}

	

	startRepl(cfg)
}