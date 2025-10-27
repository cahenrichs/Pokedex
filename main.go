package main 

import (
	"time"

	"github.com/cahenrichs/pokedexcli/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config {
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},

	}
	startRepl(cfg)
}