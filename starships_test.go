package goswapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestShip(t *testing.T) {
	sc := NewClient()

	ship, err := sc.GetShip(9)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(ship, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
