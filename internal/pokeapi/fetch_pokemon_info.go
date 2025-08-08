package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemonInfo -
func (c *Client) FetchPokemonInfo(pokemonName string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonInfo{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonInfo{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonResp := PokemonInfo{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
