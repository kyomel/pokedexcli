package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
