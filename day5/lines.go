package main

import (
	"strings"
	"unicode/utf8"
)

func niceLine(line string) bool {
	if containsThreeVowels(line) && containsDoubleLetter(line) && doesntHaveBadSubstrings(line) {
		return true
	}
	return false
}

func containsThreeVowels(line string) bool {
	vowelCount := 0

	vowels := "aeiou"
	for _, r := range line {
		if strings.ContainsRune(vowels, r) {
			vowelCount++
		}
	}

	if vowelCount >= 3 {
		return true
	}
	return false
}

func containsDoubleLetter(line string) bool {
	found := false

	for index, r := range line {
		if index == len(line)-1 {
			continue
		}
		if rune(line[index+1]) == r {
			found = true
		}
	}

	return found
}

func doesntHaveBadSubstrings(line string) bool {
	badStrings := []string{"ab", "cd", "pq", "xy"}

	for _, badString := range badStrings {
		if strings.Contains(line, badString) {
			return false
		}
	}
	return true
}

func niceSecondLine(line string) bool {
	return pairNoOverlap(line) && hasPairWithLetterInMiddle(line)
}

func pairNoOverlap(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		var pair = []byte{input[i], input[i+1]}
		for j := (i + 2); j < len(input)-1; j++ {
			if pair[0] == input[j] && pair[1] == input[j+1] {
				return true
			}
		}
	}
	return false
}

func hasPairWithLetterInMiddle(line string) bool {
	runeCount := utf8.RuneCount([]byte(line))

	for index, rune := range line {
		if index < runeCount-2 {
			twoOut, _ := utf8.DecodeRuneInString(line[index+2 : index+3])
			if twoOut == rune {
				return true
			}
		}
	}

	return false
}
