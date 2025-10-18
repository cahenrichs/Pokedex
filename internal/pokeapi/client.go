package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient gttp.Client
}

func NewClient(timeout tume.Durration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}