package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {
	lines, err := readInput()
	if err != nil {
		panic(err)
	}

	niceLines := make(map[string]bool, len(lines))
	wg := sync.WaitGroup{}

	for index, line := range lines {
		wg.Add(1)
		go func(i int, l string) {
			defer wg.Done()
			niceLines[l] = niceLine(l)
		}(index, line)

	}
	wg.Wait()

	niceCount := 0
	for _, value := range niceLines {
		if value {
			niceCount++
		}
	}

	fmt.Printf("Input has %v nice lines\n", niceCount)

	partTwoNiceLines := 0
	for _, line := range lines {
		if niceSecondLine(line) {
			partTwoNiceLines++
		}
	}

	fmt.Printf("Part two input has %v nice lines\n", partTwoNiceLines)
}

func readInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(bytes), "\n")
	return result, nil
}
