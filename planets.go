package goswapi

import (
	"strconv"
	"time"
)

//Planet struct for planets response
type Planet struct {
	Name           string        `json:"name"`
	RotationPeriod string        `json:"rotation_period"`
	OrbitalPeriod  string        `json:"orbital_period"`
	Diameter       string        `json:"diameter"`
	Climate        string        `json:"climate"`
	Gravity        string        `json:"gravity"`
	Terrain        string        `json:"terrain"`
	SurfaceWater   string        `json:"surface_water"`
	Population     string        `json:"population"`
	Residents      []interface{} `json:"residents"`
	Films          []string      `json:"films"`
	Created        time.Time     `json:"created"`
	Edited         time.Time     `json:"edited"`
	URL            string        `json:"url"`
}

//GetPlanet gets A Planet resource which is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
func (sc *SwapiClient) GetPlanet(id int) (Planet, error) {
	targetURL := baseURL + "planets/" + strconv.Itoa(id)

	p := Planet{}

	err := sc.getReq(targetURL, &p)

	return p, err
}
