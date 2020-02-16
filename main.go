package main

import (
	"fmt"

	"github.com/singularitytechnologies/openplague/internal"
)

func main() {
	fmt.Println("Starting Plague Simulation!")

	virusName := "COVID-19"

	virus := internal.NewVirus(virusName, 3.2, 0.5, 14, true, 14, true, 7, true)
	virus.CalculateTransmissionProbability(300)
	fmt.Printf("Created %s\n\n", virus.Desc())

	infectionManager := internal.NewInfectionManager(virus)
	stats := internal.NewStatistics(9e6)

	fmt.Println("Creating and Populating City")
	city := internal.NewCity(3.5e6, 1000, 1e6)
	for i := 0; i < 9e6; i++ {
		person := internal.NewPerson(infectionManager.CreateInfection, infectionManager.PersistInfection, stats.Events)
		city.AddPerson(person)
	}

	fmt.Println("Creating Patient Zero")
	joe := internal.NewPerson(infectionManager.CreateInfection, infectionManager.PersistInfection, stats.Events)
	city.AddPerson(joe)
	joe.Infect()

	for i := 0; i < 100; i++ {
		err := infectionManager.ProcessDay()

		if err != nil {
			fmt.Println("\n\nPandemic Over!!!\n\n")
			return
		}

		fmt.Printf("Day %d: %s\n", i, stats.Status())
	}
}
