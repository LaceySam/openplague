package internal

import (
	"fmt"
)

type Vehicle string

const (
	AirPlane          Vehicle = "Air Plane"
	AirPlaneOccupancy         = 600
)

type Route struct {
	Count  int
	Cities chan *City
}

type Journey struct {
	Name         string
	Kind         Vehicle
	MaxOccupancy int

	Route      *Route
	Passengers *Group
}

func NewPlaneJourney(route *Route, passengers *Group) *Journey {
	name := ""

	newCities := make(chan *City, len(r.Count))
	for city := range route.Cities {
		name += fmt.Sprintf("-%s", city.Name)
		newCities <- city
	}

	route.Cities = newCities

	return &Journey{
		Name:         name,
		Kind:         AirPlane,
		MaxOccupancy: AirPlaneOccupancy,
		Occupancy:    len(passengers.All),
		Route:        route,
		Passengers:   passengers,
	}
}

func (j *Journey) Embark(group *Group) {
	j.Passengers.Merge(group)
	return
}

func (j *Journey) Disembark(count int) *Group {
	return j.Passengers.ExtractGroup(count)
}

func (j *Journey) AvailableSpace() int {
	return j.MaxOccupancy - len(j.Passengers.All)
}

func (j *Journey) Proccess() {
	i := 1
	for city := range j.Route.Cities {
		var count int
		if j.Route.Count < i {
			count := random.Intn(len(j.Passengers.All))
		} else {
			count := len(j.Passengers.All)
		}

		// Let a random set off at destination
		disembarkers := j.Disembark(count)
		for _, person := range disembarkers.All {
			city.AddPerson(person)
		}

		if j.Route.Count < i {
			// Add new people onto journey
			newCount := random.Intn(j.AvailableSpace())
			embarkers := city.ExtractGroup(j.Passengers.SickEncounters, j.Passengers.MaxEncounters, newCount)
			j.Embark(embarkers)
		}
	}
}

type TransportationHub interface {
	Arrivals() int
	Departures() int
	MaxOccupancy() int
	AverageOccupancy() int
	CreateJourney(Group) *Journey
	AddRoute(*Route) bool
	DeleteRoute(*Route) bool
}

type Airport struct {
}
