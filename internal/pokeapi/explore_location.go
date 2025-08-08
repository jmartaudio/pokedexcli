package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocationExplore -
func (c *Client) ExploreLocations(locationName string) (LocationExplore, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationExplore{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationExplore{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationExplore{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationExplore{}, err
	}

	locationResp := LocationExplore{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationExplore{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
