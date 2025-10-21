package internal

import (

)

func (c *Client) GetLocation() {
	url := baseURL + "/loaction-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			reutrn Location{}, err
		}
		return locationResp, nil
	}
}