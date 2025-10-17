package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

func getLocation() {
	url := baseURL + "/location-area/" 

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}