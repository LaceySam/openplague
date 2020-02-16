package internal

import (
	"fmt"
)

var (
	DEATH        = "death"
	SICK         = "sick"
	RECOVERED    = "recovered"
	INFECTED     = "infected"
	CONTAGIOUS   = "contagious"
	UNCONTAGIOUS = "uncontagious"
)

type Event struct {
	Name string
	Kind string
}

func NewEvent(name, kind string) *Event {
	return &Event{Name: name, Kind: kind}
}

type Statistics struct {
	Susceptible int
	Infected    int
	Immune      int
	Recovered   int
	Contagious  int
	Sick        int
	Deaths      int
}

func NewStatistics(population int) *Statistics {
	return &Statistics{
		Susceptible: population,
		Infected:    0,
		Immune:      0,
		Recovered:   0,
		Contagious:  0,
		Sick:        0,
		Deaths:      0,
	}
}

func (s *Statistics) Events(event *Event) {
	switch event.Kind {
	case DEATH:
		s.Infected--
		s.Deaths++
		s.Sick--
		s.Contagious--
	case RECOVERED:
		s.Infected--
		s.Recovered++
		s.Immune++
		s.Sick--
	case INFECTED:
		s.Infected++
		s.Susceptible--
	case CONTAGIOUS:
		s.Contagious++
	case UNCONTAGIOUS:
		s.Contagious--
	case SICK:
		s.Sick++
	}
}

func (s *Statistics) Status() string {
	return fmt.Sprintf(
		"Susceptible: %d, Infected %d, Recovered %d, Contagious: %d, Sick: %d, Deaths: %d",
		s.Susceptible,
		s.Infected,
		s.Recovered,
		s.Contagious,
		s.Sick,
		s.Deaths,
	)
}
