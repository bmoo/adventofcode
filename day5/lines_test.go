package main

import "testing"

func TestNiceLineNoDoubleLetter(t *testing.T) {
	line := "jchzalrnumimnmhp"
	nice := niceLine(line)

	if nice {
		t.Errorf("%v should not be a nice line", line)
	}
}

func TestNiceLineThreeVowels(t *testing.T) {
	line := "ugknbfddgicrmopn"
	nice := niceLine(line)

	if !nice {
		t.Errorf("%v should be a nice line", line)
	}
}

func TestThreeVowelsAndDoubleLetter(t *testing.T) {
	line := "aaa"
	nice := niceLine(line)

	if !nice {
		t.Errorf("%v should be a nice line", line)
	}
}

func TestContainsXY(t *testing.T) {
	line := "haegwjzuvuyypxyu"
	nice := niceLine(line)

	if nice {
		t.Errorf("%v should not be a nice line", line)
	}
}

func TestContainsOneVowel(t *testing.T) {
	line := "dvszwmarrgswjxmb"
	nice := niceLine(line)

	if nice {
		t.Errorf("%v should not be a nice line", line)
	}
}
