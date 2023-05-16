package main

import (
	"github.com/bigveezus/my-cli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}

	

	startRepl(cfg)
}