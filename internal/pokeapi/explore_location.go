package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Explore Location -
func (c *Client) ExploreLocations(area *string) (LocationInfo, error) {
	url := baseURL + "/location-area/" + *area
	fmt.Printf("Exploring... %s\n\n", url)
	if *area == "" {
		fmt.Println("Enter Location to Expolre")
	}

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationInfo{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationInfo{}, err
		}
		defer resp.Body.Close()
		
		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationInfo{}, err
		}

		c.cache.Add(url, dat)

		locExpResp := LocationInfo{}
		err = json.Unmarshal(dat, &locExpResp)
		if err != nil {
			return LocationInfo{}, err
		}

		return locExpResp, nil
	} else {
		locExpResp := LocationInfo{}
		err := json.Unmarshal(dat, &locExpResp)
		if err != nil {
			return LocationInfo{}, err
	}

	return locExpResp, nil
	}
}
