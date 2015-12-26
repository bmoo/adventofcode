package main

import (
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	dim, err := readLine("5x11x12")
	if err != nil {
		t.Error(err)
	}

	if dim.length != 5 {
		t.Errorf("Expected length 5 but got %v", dim.length)
	}

	if dim.width != 11 {
		t.Errorf("Expected width 11 but got %v", dim.width)
	}

	if dim.height != 12 {
		t.Errorf("Expected height 12 but got %v", dim.height)
	}
}

func TestReadInput(t *testing.T) {
	dims, err := readInput(strings.NewReader("5x11x12\n4x2x3"))
	if err != nil {
		t.Error(err)
	}

	if len(dims) != 2 {
		t.Errorf("Expected dimensions of length 2 but got %v", len(dims))
	}
}

func TestExtraSlack(t *testing.T) {
	allDims := AllDimensions{&Dimensions{length: 1, width: 1, height: 10}}
	if allDims.totalArea() != 43 {
		t.Errorf("Expected 43 feet but got %v", allDims.totalArea())
	}

	allDims = AllDimensions{&Dimensions{length: 2, width: 3, height: 4}}
	if allDims.totalArea() != 58 {
		t.Errorf("Expected 58 feet but got %v", allDims.totalArea())
	}
}

func TestRibbonLength(t *testing.T) {
	allDims := AllDimensions{&Dimensions{length: 1, width: 1, height: 10}}
	if allDims.totalRibbonLength() != 14 {
		t.Errorf("Expected 14 feet but got %v", allDims.totalRibbonLength())
	}

	allDims = AllDimensions{&Dimensions{length: 2, width: 3, height: 4}}
	if allDims.totalRibbonLength() != 34 {
		t.Errorf("Expected 34 feet but got %v", allDims.totalRibbonLength())
	}
}

func TestSmallestPerimeter(t *testing.T) {
	dims := &Dimensions{length: 1, width: 1, height: 10}
	if dims.ribbonLength() != 4 {
		t.Errorf("Expected 4 feet of ribbon but got %v", dims.ribbonLength())
	}

}
