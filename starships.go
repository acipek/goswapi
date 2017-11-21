package goswapi

import (
	"strconv"
	"time"
)

//Starship struct for starships response
type Starship struct {
	Name                 string        `json:"name"`
	Model                string        `json:"model"`
	Manufacturer         string        `json:"manufacturer"`
	CostInCredits        string        `json:"cost_in_credits"`
	Length               string        `json:"length"`
	MaxAtmospheringSpeed string        `json:"max_atmosphering_speed"`
	Crew                 string        `json:"crew"`
	Passengers           string        `json:"passengers"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	HyperdriveRating     string        `json:"hyperdrive_rating"`
	MGLT                 string        `json:"MGLT"`
	StarshipClass        string        `json:"starship_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              time.Time     `json:"created"`
	Edited               time.Time     `json:"edited"`
	URL                  string        `json:"url"`
}

//GetShip gets a Starship resource which is a single transport craft that has hyperdrive capability.
func (sc *SwapiClient) GetShip(id int) (Starship, error) {
	targetURL := baseURL + "starships/" + strconv.Itoa(id)

	s := Starship{}

	err := sc.getReq(targetURL, &s)

	return s, err
}
