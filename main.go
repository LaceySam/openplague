package main

import (
	"fmt"
	"time"

	"github.com/singularitytechnologies/openplague/internal"
)

func main() {
	fmt.Println("Starting Plague Simulation!")

	virusName := "COVID-19"

	virus := internal.NewVirus(virusName, 3.2, 0.03, 14, true, 14, true, 7, true)
	virus.CalculateTransmissionProbability(300)
	fmt.Printf("Created %s\n\n", virus.Desc())

	infectionManager := internal.NewInfectionManager(virus)
	population := 8900000
	stats := internal.NewStatistics(population)

	fmt.Println("Creating and Populating City")
	city := internal.NewCity(3.5e6, 1000, 1e6)
	for i := 0; i < population; i++ {
		person := internal.NewPerson(infectionManager.CreateInfection, infectionManager.PersistInfection, stats.Events)
		eventFns := city.AddPerson(person)
		person.AddEventFns(eventFns)
	}

	fmt.Println("Creating Patient Zero")
	joe := internal.NewPerson(infectionManager.CreateInfection, infectionManager.PersistInfection, stats.Events)
	eventFns := city.AddPerson(joe)
	joe.AddEventFns(eventFns)
	joe.Infect()

	i := 0
	for {
		time.Sleep(time.Millisecond)
		infectionManager.ProcessDay()
		city.ProcessDay()

		fmt.Printf("Day %d: %s\n", i, stats.Status())
		i++

		if stats.Infected == 0 {
			break
		}
	}

	fmt.Println("\n\n\nPandemic Over!!!\n\n\n")
}
