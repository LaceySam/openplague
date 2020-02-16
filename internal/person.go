package internal

import (
	"fmt"

	"github.com/google/uuid"
)

type UpdateInfection func(infection *Infection)
type CreateInfection func(updateInfection UpdateInfection)

type Person struct {
	Name     string
	Infected bool
	// Whether a person can be infected
	Immune bool
	// Whether a person can infect another
	Contagious bool
	// Whether or not person stays at home or not
	Active bool
	Alive  bool

	createInfection CreateInfection
}

func NewPerson(createInfection CreateInfection) *Person {
	return &Person{
		Name:            uuid.New().String(),
		Active:          true,
		Alive:           true,
		createInfection: createInfection,
	}
}

func (p *Person) Infect() {
	if !p.Infected && !p.Immune {
		p.createInfection(p.UpdateInfection)
		p.Infected = true
		// TODO(Sam): Broadcast events
	}
}

func (p *Person) Kill() {
	fmt.Println("RIP")
	p.Alive = false

	// TODO(Sam): Broadcast events
	panic("RIP")
}

func (p *Person) UpdateInfection(infection *Infection) {
	p.Contagious = infection.Contagious()
	if infection.Active() || infection.Recovery() {
		p.Active = false
	}

	if infection.Active() && infection.KillPatient() {
		p.Kill()
	}

	if infection.Recovery() {
		p.Immune = true
	}

	if infection.Complete() {
		p.Infected = false
		p.Active = true
	}
}

func (p *Person) Status() string {
	return fmt.Sprintf(
		"Person %q: Infected %t, Contagious %t, Immune %t, Active %t, Alive %t",
		p.Name,
		p.Infected,
		p.Contagious,
		p.Immune,
		p.Active,
		p.Alive,
	)
}
