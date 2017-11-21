package goswapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSpecies(t *testing.T) {
	sc := NewClient()

	species, err := sc.GetSpecies(1)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(species, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
