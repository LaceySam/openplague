package internal

import (
	"fmt"

	"github.com/google/uuid"
)

type UpdateInfection func(infection *Infection)
type CreateInfection func(updateInfection UpdateInfection)
type EventFn func(event *Event)

type Person struct {
	Name     string
	Infected bool
	// Whether a person can be infected
	Immune bool
	// Whether a person can infect another
	Contagious bool
	// Whether or not person stays at home or not
	Sick  bool
	Alive bool

	createInfection  CreateInfection
	persistInfection UpdateInfection
	events           EventFn

	Infection *Infection
}

func NewPerson(createInfection CreateInfection, persistInfection UpdateInfection, events EventFn) *Person {
	return &Person{
		Name:             uuid.New().String(),
		Alive:            true,
		createInfection:  createInfection,
		persistInfection: persistInfection,
		events:           events,
	}
}

func (p *Person) Infect() {
	if !p.Infected && !p.Immune {
		p.createInfection(p.UpdateInfection)
		p.Infected = true
		p.events(NewEvent(p.Name, INFECTED))
	}
}

func (p *Person) Kill() {
	p.Alive = false
	p.events(NewEvent(p.Name, DEATH))
}

func (p *Person) UpdateInfection(infection *Infection) {
	p.Infection = infection

	if !p.Contagious && infection.Contagious() {
		p.Contagious = true
		p.events(NewEvent(p.Name, CONTAGIOUS))
	}

	if p.Contagious && !infection.Contagious() {
		p.Contagious = false
		p.events(NewEvent(p.Name, UNCONTAGIOUS))
	}

	p.Contagious = infection.Contagious()

	if !p.Sick && infection.Active() {
		p.Sick = true
		p.events(NewEvent(p.Name, SICK))
	}

	if infection.Active() && infection.KillPatient() {
		p.Kill()
		return
	}

	if infection.Recovery() {
		p.Immune = true
	}

	if infection.Complete() {
		p.Infected = false
		p.Sick = false
		p.events(NewEvent(p.Name, RECOVERED))
		p.Infection = nil
		return
	}

	p.persistInfection(infection)
}

func (p *Person) Status() string {
	return fmt.Sprintf(
		"Person %q: Infected %t, Contagious %t, Immune %t, Sick %t, Alive %t",
		p.Name,
		p.Infected,
		p.Contagious,
		p.Immune,
		p.Sick,
		p.Alive,
	)
}
