package goswapi

import (
	"strconv"
	"time"
)

//Vehicle struct for vehicles response
type Vehicle struct {
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
	VehicleClass         string        `json:"vehicle_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              time.Time     `json:"created"`
	Edited               time.Time     `json:"edited"`
	URL                  string        `json:"url"`
}

//GetVehicle gets a Vehicle resource which is a single transport craft that does not have hyperdrive capability.
func (sc *SwapiClient) GetVehicle(id int) (Vehicle, error) {
	targetURL := baseURL + "vehicles/" + strconv.Itoa(id)

	v := Vehicle{}

	err := sc.getReq(targetURL, &v)

	return v, err
}
