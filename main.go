package main 

import (
	"time"

	"github.com/cahenrichs/pokedexcli/internal/pokeapi"
)


func main() {
	cfg := &config {
		pokeapiClient: pokeClient,

	}
	startRepl(cfg)
}