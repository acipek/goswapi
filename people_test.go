package goswapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPeople(t *testing.T) {
	sc := NewClient()

	people, err := sc.GetPeople(1)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
