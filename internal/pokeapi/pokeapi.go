package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/atishay-aj/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type SingleAreaResp struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	// check cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		locationAreaResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad req %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}

func (c *Client) GetSingleArea(areaName string) (SingleAreaResp, error) {
	endpoint := "/location-area/" + areaName
	fullURL := baseURL + endpoint

	// check cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		singleAreaResp := SingleAreaResp{}
		err := json.Unmarshal(data, &singleAreaResp)
		if err != nil {
			return SingleAreaResp{}, err
		}
		return singleAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return SingleAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SingleAreaResp{}, err
	}
	if resp.StatusCode > 399 {
		return SingleAreaResp{}, fmt.Errorf("bad req %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return SingleAreaResp{}, err
	}
	singleAreaResp := SingleAreaResp{}
	err = json.Unmarshal(data, &singleAreaResp)
	if err != nil {
		return SingleAreaResp{}, err
	}
	c.cache.Add(fullURL, data)
	return singleAreaResp, nil
}
