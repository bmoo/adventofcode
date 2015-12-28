package main

import "fmt"

type location struct {
	x int
	y int
}

type Santa struct {
	currentLocation location
	Stops           map[location]int
}

func NewSanta() *Santa {
	santa := &Santa{}
	santa.Stops = make(map[location]int)
	santa.Stops[santa.currentLocation] = 1

	return santa
}

func (h location) hash() int {
	return (h.x * 31) + (h.y * 17)
}

func (s *Santa) MoveToNewLocation(code rune) error {
	l := s.currentLocation

	switch code {
	case '^':
		s.currentLocation = location{x: l.x, y: l.y + 1}
	case '>':
		s.currentLocation = location{x: l.x + 1, y: l.y}
	case '<':
		s.currentLocation = location{x: l.x - 1, y: l.y}
	case 'v':
		s.currentLocation = location{x: l.x, y: l.y - 1}
	default:
		return fmt.Errorf("Expected a valid location code but got %v", code)
	}

	val, exists := s.Stops[s.currentLocation]
	if exists {
		s.Stops[s.currentLocation] = val + 1
	} else {
		s.Stops[s.currentLocation] = 1
	}

	return nil
}
