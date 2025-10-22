package pokeapi

import (
	"net/http"
	"time"

	"github.com/cahenrichs/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache	*pokecache.Cache
}

// go
func NewClient(timeout time.Duration) Client {
    return Client{
        httpClient: http.Client{ Timeout: timeout },
        cache:      pokecache.NewCache(5 * time.Minute),
    }
}