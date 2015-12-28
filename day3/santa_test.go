package main

import "testing"

func TestParseStops(t *testing.T) {
	var input string = "^v^v^v^v^v"
	santa, err := parseSantaStops(input)
	if err != nil {
		t.Error(err)
	}

	if len(santa.Stops) != 2 {
		t.Errorf("Expected 2 stops but got %v", len(santa.Stops))
	}
}

func TestMoreStops(t *testing.T) {
	input := "^>v<"
	santa, err := parseSantaStops(input)
	if err != nil {
		t.Error(err)
	}

	if len(santa.Stops) != 4 {
		t.Errorf("Expected 4 stops but got %v", len(santa.Stops))
	}
}

func TestCountFirstStop(t *testing.T) {
	input := ">"
	santa, err := parseSantaStops(input)
	if err != nil {
		t.Error(err)
	}

	if len(santa.Stops) != 2 {
		t.Errorf("Expected 2 stops but got %v", len(santa.Stops))
	}
}
