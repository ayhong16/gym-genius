package src

import (
	"encoding/json"
	"net/http"
	"server/initializers"
	"server/types"
)

type API struct {
	apiKey string
}

func NewAPI() *API {
	apiKey := initializers.LoadAPIKey()
	return &API{apiKey: apiKey}
}

func (api *API) FetchExercises() ([]types.Exercise, error) {
	url := "https://exercisedb.p.rapidapi.com/exercises?limit=0&offset=0"

	resp, err := api.getURL(url)
	if err != nil {
		return nil, err
	}

	var exercises []types.Exercise
	if err := json.NewDecoder(resp.Body).Decode(&exercises); err != nil {
		return nil, err
	}

	return exercises, nil
}

func (api *API) FetchBodyParts() (string, error) {
	url := "https://exercisedb.p.rapidapi.com/bodyPartList"

	resp, err := api.getURL(url)
	if err != nil {
		return "", err
	}

	var bodyParts string
	if err := json.NewDecoder(resp.Body).Decode(&bodyParts); err != nil {
		return "", err
	}
	return bodyParts, nil
}

func (api *API) getURL(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Add("x-rapidapi-host", "exercisedb.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", api.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &http.Response{}, err
	}
	return resp, nil
}
