package goswapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPlanet(t *testing.T) {
	sc := NewClient()

	planet, err := sc.GetPlanet(1)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(planet, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
