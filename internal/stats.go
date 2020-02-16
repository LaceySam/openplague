package internal

import (
	"fmt"
)

var (
	DEATH     Event = "death"
	RECOVERED Event = "recovered"
	INFECTED  Event = "infected"
)

type Event string

type Statistics struct {
	Susceptible int
	Infected    int
	Immune      int
	Recovered   int
	Deaths      int
}

func NewStatistics(population int) *Statistics {
	return &Statistics{
		Susceptible: population,
		Infected:    0,
		Immune:      0,
		Recovered:   0,
		Deaths:      0,
	}
}

func (s *Statistics) Events(event Event) {
	switch event {
	case DEATH:
		s.Infected--
		s.Deaths++
	case RECOVERED:
		s.Infected--
		s.Recovered++
		s.Immune++
	case INFECTED:
		s.Infected++
		s.Susceptible--
	}
}

func (s *Statistics) Status() string {
	return fmt.Sprintf(
		"Susceptible: %d, Infected %d, Recovered %d, Deaths %d",
		s.Susceptible,
		s.Infected,
		s.Recovered,
		s.Deaths,
	)
}
