package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(locationAreaName *string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + *locationAreaName

	if val, ok := c.cache.Get(url); ok {
		locationsAreaRes := RespLocationArea{}
		err := json.Unmarshal(val, &locationsAreaRes)
		if err != nil {
			return RespLocationArea{}, err
		}
		return locationsAreaRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationArea{}, err
	}
	locationsRes := RespLocationArea{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, data)
	return locationsRes, nil
}
