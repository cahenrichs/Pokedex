package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchLocationAreas(url string) (LocationAreaList, error) {
	url := baseURL + "/location-area" + 
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaList{}, err
	}

	defer.Body.Close()
	
	if resp.Statusode != http.StatusOK {
		return LocationAreaList{}, fmt.Errorf("bad status %d", resp.StatusCode)
	}

	data,err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	var out LocationAreaList
	if err := json.Unmarshal(data, &out); err != nil {
		return LocationAreaList{}, err
	}
	return out, nil
}