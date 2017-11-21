package goswapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://swapi.co/api/"

//SwapiClient client for swapi
type SwapiClient struct {
	HTTPClient *http.Client
}

//NewClient creates new swapi client
func NewClient() SwapiClient {
	return SwapiClient{
		HTTPClient: &http.Client{},
	}
}

func (sc *SwapiClient) getReq(url string, out interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request")
	}

	res, err := sc.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %v", res.Status)
	}

	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}
