package main

import (
	"testing"
)

func TestParseFloor(t *testing.T) {
	input := "((())"

	result := parseFloor(input)
	if result.FinalFloor != 1 {
		t.Errorf("Expected 1 but got %v", result.FinalFloor)
	}
}

func TestFirstBasement(t *testing.T) {
	input := ")())))))"

	result := parseFloor(input)
	if result.FirstBasementPosition == nil {
		t.Error("Expected non-nil first basement position")
	}

	if *result.FirstBasementPosition != 1 {
		t.Errorf("Expected 1 but got %v", *result.FirstBasementPosition)
	}
}
