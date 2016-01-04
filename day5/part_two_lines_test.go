package main

import "testing"

func TestFirst(t *testing.T) {
	line := "qjhvhtzxzqqjkmpb"

	nice := niceSecondLine(line)
	if !nice {
		t.Errorf("%v should be nice", line)
	}
}

func TestSecond(t *testing.T) {
	line := "xxyxx"

	nice := niceSecondLine(line)
	if !nice {
		t.Errorf("%v should be nice", line)
	}
}

func TestThird(t *testing.T) {
	line := "uurcxstgmygtbstg"

	nice := niceSecondLine(line)
	if nice {
		t.Errorf("%v should be naughty", line)
	}
}

func TestFourth(t *testing.T) {
	line := "ieodomkazucvgmuy"

	nice := niceSecondLine(line)
	if nice {
		t.Errorf("%v should be naughty", line)
	}
}

func TestTwoPairsInLine(t *testing.T) {
	line := "aabcdefgaa"
	twoPairs := pairNoOverlap(line)

	if !twoPairs {
		t.Errorf("%v should have two pairs in it", line)
	}
}

func TestHasPairWithLetterInMiddle(t *testing.T) {
	line := "abcdefeghi"
	hasPair := hasPairWithLetterInMiddle(line)

	if !hasPair {
		t.Errorf("%v should have a pair separated by a letter", line)
	}
}

func TestPairDoesntOverlap(t *testing.T) {
	line := "aaa"
	nice := pairNoOverlap(line)

	if nice {
		t.Errorf("%v should not be nice because it overlaps", line)
	}
}

func TestHasPairInLineWithBad(t *testing.T) {
	line := "ieodomkazucvgmuy"
	nice := pairNoOverlap(line)

	if nice {
		t.Errorf("There shouldn't be a pair in %v", line)
	}
}

func TestAdditionalValue(t *testing.T) {
	line := "xdwduffwgcptfwad"
	nice := niceSecondLine(line)

	if !nice {
		t.Errorf("Line %v should be nice", line)
	}
}
