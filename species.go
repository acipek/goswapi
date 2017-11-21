package goswapi

import (
	"strconv"
	"time"
)

//Species struct for species response
type Species struct {
	Name            string    `json:"name"`
	Classification  string    `json:"classification"`
	Designation     string    `json:"designation"`
	AverageHeight   string    `json:"average_height"`
	SkinColors      string    `json:"skin_colors"`
	HairColors      string    `json:"hair_colors"`
	EyeColors       string    `json:"eye_colors"`
	AverageLifespan string    `json:"average_lifespan"`
	Homeworld       string    `json:"homeworld"`
	Language        string    `json:"language"`
	People          []string  `json:"people"`
	Films           []string  `json:"films"`
	Created         time.Time `json:"created"`
	Edited          time.Time `json:"edited"`
	URL             string    `json:"url"`
}

//GetSpecies gets a Species resource which is a type of person or character within the Star Wars Universe.
func (sc *SwapiClient) GetSpecies(id int) (Species, error) {
	targetURL := baseURL + "species/" + strconv.Itoa(id)

	s := Species{}

	err := sc.getReq(targetURL, &s)

	return s, err
}
