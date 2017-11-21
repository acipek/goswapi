package goswapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVehicle(t *testing.T) {
	sc := NewClient()

	vehicle, err := sc.GetVehicle(4)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(vehicle, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
