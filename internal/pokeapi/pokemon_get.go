package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}


	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
 	   return Pokemon{}, err
	}

	var p Pokemon
	if err := json.Unmarshal(data, &p); err != nil {
    	return Pokemon{}, err
	}

	return p, nil

}