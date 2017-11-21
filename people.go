package goswapi

import (
	"strconv"
	"time"
)

//People struct for people response
type People struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

//GetPeople gets person or character within the Star Wars universe by id
func (sc *SwapiClient) GetPeople(id int) (People, error) {
	targetURL := baseURL + "people/" + strconv.Itoa(id)

	p := People{}

	err := sc.getReq(targetURL, &p)

	return p, err
}
