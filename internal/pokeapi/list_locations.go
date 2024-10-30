package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsRes := RespLocations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return RespLocations{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}
	locationsRes := RespLocations{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, data)
	return locationsRes, nil
}
