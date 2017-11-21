package goswapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFilms(t *testing.T) {
	sc := NewClient()

	film, err := sc.GetFilm(1)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(film, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
