package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()
		
		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, dat)

		locationsResp := RespShallowLocations{}
		err = json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	} else {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
	}

	return locationsResp, nil
	}
}
