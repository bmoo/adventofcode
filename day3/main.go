package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	santa, err := parseSantaStops(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Santa made %v stops\n", len(santa.Stops))

	santa, robosanta, err := parseSantaAndRoboSantaStops(input)
	if err != nil {
		panic(err)
	}

	combinedStops := appendStops(santa, robosanta)
	fmt.Printf("RoboSanta and santa made %v different stops\n", len(combinedStops))
}

func parseSantaStops(input string) (*Santa, error) {
	santa := NewSanta()

	for _, char := range input {
		err := santa.MoveToNewLocation(char)
		if err != nil {
			return nil, err
		}
	}

	return santa, nil
}

func parseSantaAndRoboSantaStops(input string) (*Santa, *Santa, error) {
	santa := NewSanta()
	robosanta := NewSanta()

	for index, char := range input {
		if index%2 == 0 {
			err := santa.MoveToNewLocation(char)
			if err != nil {
				return nil, nil, err
			}
		} else {
			err := robosanta.MoveToNewLocation(char)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return santa, robosanta, nil
}

func appendStops(santa *Santa, robosanta *Santa) map[location]int {
	result := make(map[location]int)

	for k, v := range santa.Stops {
		result[k] = v
	}

	for k, v := range robosanta.Stops {
		value, exists := result[k]
		if exists {
			result[k] = value + v
		} else {
			result[k] = v
		}
	}

	return result
}

func readInput() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", nil
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}
