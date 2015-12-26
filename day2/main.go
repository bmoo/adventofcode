package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	dims, err := readInput(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %v dimensions from the input\n", len(dims))

	fmt.Println("The total area in the file is", dims.totalArea())

	fmt.Println("The total ribbon needed is", dims.totalRibbonLength())
}

type Dimensions struct {
	length int
	width  int
	height int
}

type AllDimensions []*Dimensions

func readInput(reader io.Reader) (AllDimensions, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	result := make([]*Dimensions, 0)

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		dim, err := readLine(line)
		if err != nil {
			return nil, err
		}

		if dim != nil {
			result = append(result, dim)
		}
	}

	return AllDimensions(result), nil
}

func readLine(line string) (*Dimensions, error) {
	// is this just whitespace?
	if len(strings.TrimSpace(line)) == 0 {
		return nil, nil
	}

	dimensions := strings.Split(line, "x")

	if len(dimensions) != 3 {
		return nil, errors.New(fmt.Sprintf("Line does not have length of 3. Was ", len(dimensions)))
	}

	dim := &Dimensions{}
	var err error
	dim.length, err = strconv.Atoi(dimensions[0])
	if err != nil {
		return nil, errors.New("Could not parse length")
	}
	dim.width, err = strconv.Atoi(dimensions[1])
	if err != nil {
		return nil, errors.New("Could not parse width")
	}
	dim.height, err = strconv.Atoi(dimensions[2])
	if err != nil {
		return nil, errors.New("Could not parse height")
	}

	return dim, nil
}

func (dims AllDimensions) totalArea() int {
	var result int = 0

	for _, dim := range dims {
		result += dim.area()
	}

	return result
}

func (dim *Dimensions) area() int {
	sides := dim.sidesArray()

	min := sides[0]

	total := 0
	for _, value := range sides {
		total += (2 * value)
		if value < min {
			min = value
		}
	}

	return total + min
}

func (dims AllDimensions) totalRibbonLength() int {
	var result int = 0
	for _, dim := range dims {
		result += dim.bowLength() + dim.ribbonLength()
	}

	return result
}

func (dim *Dimensions) sidesArray() []int {
	sides := make([]int, 3)

	sides[0] = dim.length * dim.width
	sides[1] = dim.width * dim.height
	sides[2] = dim.height * dim.length

	return sides
}

func (dim *Dimensions) ribbonLength() int {
	sides := dim.edgesArray()

	minIndex := 0
	for index, value := range sides {
		if value < sides[minIndex] {
			minIndex = index
		}
	}

	secondMinIndex := 0
	if minIndex == 0 {
		secondMinIndex = 1
	}

	for index, value := range sides {
		if index == minIndex {
			continue
		}

		if value < sides[secondMinIndex] {
			secondMinIndex = index
		}
	}

	return (2 * sides[minIndex]) + (2 * sides[secondMinIndex])
}

func (dim *Dimensions) edgesArray() []int {
	sides := make([]int, 3)

	sides[0] = dim.length
	sides[1] = dim.width
	sides[2] = dim.height

	return sides
}

func (dim *Dimensions) bowLength() int {
	return dim.width * dim.height * dim.length
}
