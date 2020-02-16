package main

import (
	"fmt"

	"github.com/singularitytechnologies/openplague/internal"
)

func main() {
	fmt.Println("Starting Plague Simulation!")

	virusName := "COVID-19"

	virus := internal.NewVirus(virusName, 3.2, 0.03, 14, true, 14, true, 7, true)
	fmt.Printf("Created %s\n", virus.Desc())

	infectionManager := internal.NewInfectionManager(virus)

	joe := internal.NewPerson(infectionManager.CreateInfection)
	joe.Infect()
	fmt.Println(joe.Status())

	for i := 0; i < 100; i++ {
		fmt.Printf("Processing Day %d\n", i)
		err := infectionManager.ProcessDay()
		if err != nil {
			fmt.Println("\n\nPandemic Over!!!\n\n")
			return
		}

		fmt.Println(joe.Status())
	}
}
