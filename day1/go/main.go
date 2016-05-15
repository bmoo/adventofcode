package main

import (
	"fmt"
	"io/ioutil"
)

type ParseOutput struct {
	FinalFloor            int
	FirstBasementPosition *int
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := parseFloor(string(bytes))

	fmt.Printf("The result is %v\n", result.FinalFloor)
	fmt.Printf("Enters the basement at %v\n", *result.FirstBasementPosition)
}

func parseFloor(clues string) ParseOutput {
	var result ParseOutput = ParseOutput{}
	for index, clue := range clues {
		if clue == '(' {
			result.FinalFloor++
		}
		if clue == ')' {
			result.FinalFloor--
		}
		if result.FirstBasementPosition == nil {
			if result.FinalFloor < 0 {
				firstPosition := index + 1
				result.FirstBasementPosition = &firstPosition
			}
		}

	}
	return result
}
