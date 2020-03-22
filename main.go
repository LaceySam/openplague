package main

import (
	"fmt"
	"time"

	"github.com/singularitytechnologies/openplague/internal"
)

func main() {
	fmt.Println("Starting Pandemic Simulation!")

	virus := internal.NewVirus("COVID-19", 3.2, 0.033, 14, true, 14, true, 7, true)
	//virus := internal.NewVirus("Swine Flu", 1.6, 0.0002, 4, true, 7, true, 7, false)
	virus.CalculateTransmissionProbability(100)
	fmt.Printf("Created %s\n\n", virus.Desc())

	infectionManager := internal.NewInfectionManager(virus)
	//population := 9000000
	population := 1000
	stats := internal.NewStatistics(population)

	fmt.Println("Creating and Populating City")
	city := internal.NewCity("London", 3.5e6, 1000, 1e6)
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
