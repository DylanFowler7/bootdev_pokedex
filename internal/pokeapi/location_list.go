package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetLocation(locationAreaName string) (LocationEncounters, error) {
	endpoint := "/location-area/" + locationAreaName
	locationUrl := baseURL + endpoint

	dat, ok := c.cache.Get(locationUrl)
	if ok {
		fmt.Println("cache hit!")
		locationEncounters := LocationEncounters{}
		err := json.Unmarshal(dat, &locationEncounters)
		if err != nil {
			return LocationEncounters{}, err
		}
		return locationEncounters, nil
	}
	fmt.Println("cache miss!")
	req, err := http.NewRequest("GET", locationUrl, nil)
	if err != nil {
		return LocationEncounters{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationEncounters{}, err
	}
	defer resp.Body.Close()
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationEncounters{}, err
	}
	locationEncounters := LocationEncounters{}
	err = json.Unmarshal(dat, &locationEncounters)
	if err != nil {
		return LocationEncounters{}, err
	}
	c.cache.Add(locationUrl, dat)
	return locationEncounters, nil
}

// ListLocations -
func (c *Client) GetLocationArea(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationArea, nil
}

