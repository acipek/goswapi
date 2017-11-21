package goswapi

import (
	"strconv"
	"time"
)

//Film struct for films response
type Film struct {
	Title        string    `json:"title"`
	EpisodeID    int       `json:"episode_id"`
	OpeningCrawl string    `json:"opening_crawl"`
	Director     string    `json:"director"`
	Producer     string    `json:"producer"`
	ReleaseDate  string    `json:"release_date"`
	Characters   []string  `json:"characters"`
	Planets      []string  `json:"planets"`
	Starships    []string  `json:"starships"`
	Vehicles     []string  `json:"vehicles"`
	Species      []string  `json:"species"`
	Created      time.Time `json:"created"`
	Edited       time.Time `json:"edited"`
	URL          string    `json:"url"`
}

//GetFilm gets a Film resource by id
func (sc *SwapiClient) GetFilm(id int) (Film, error) {
	targetURL := baseURL + "films/" + strconv.Itoa(id)

	f := Film{}

	err := sc.getReq(targetURL, &f)

	return f, err
}
