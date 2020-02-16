package internal

import (
	"math/rand"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

type Infection struct {
	Name                    string
	MortalityProbability    float64
	IncubationPeriod        int
	InfectiousIncubation    bool
	ActivePeriod            int
	InfectiousActive        bool
	RecoveryPeriod          int
	InfectiousRecovery      bool
	UpdateInfection         UpdateInfection
	day                     int
	TransmissionProbability float64
}

func NewInfection(virus *Virus, updateInfection UpdateInfection) *Infection {
	// TODO(Sam): Slightly change values per infection here
	return &Infection{
		Name:                    virus.Name,
		MortalityProbability:    virus.MortalityProbability,
		IncubationPeriod:        virus.IncubationPeriod,
		InfectiousIncubation:    virus.InfectiousIncubation,
		ActivePeriod:            virus.ActivePeriod,
		InfectiousActive:        virus.InfectiousActive,
		RecoveryPeriod:          virus.RecoveryPeriod,
		InfectiousRecovery:      virus.InfectiousRecovery,
		TransmissionProbability: virus.TransmissionProbability,
		UpdateInfection:         updateInfection,
	}
}

func (i *Infection) Incubating() bool {
	if i.day < i.IncubationPeriod {
		return true
	}

	return false
}

func (i *Infection) Active() bool {
	if !i.Incubating() && i.day < i.IncubationPeriod+i.ActivePeriod {
		return true
	}

	return false
}

func (i *Infection) Recovery() bool {
	if !i.Incubating() && !i.Active() && i.day < i.IncubationPeriod+i.ActivePeriod+i.RecoveryPeriod {
		return true
	}

	return false
}

func (i *Infection) Complete() bool {
	if !i.Incubating() && !i.Active() && !i.Recovery() {
		return true
	}

	return false
}

func (i *Infection) Contagious() bool {
	if i.Incubating() && i.InfectiousIncubation {
		return true
	}

	if i.Active() && i.InfectiousActive {
		return true
	}

	if i.Recovery() && i.InfectiousRecovery {
		return true
	}

	return false
}

func (i *Infection) KillPatient() bool {
	if !i.Active() {
		return false
	}

	if random.Float64() < i.MortalityProbability {
		return true
	}

	return false
}

func (i *Infection) Progress() bool {
	i.day++
	i.UpdateInfection(i)
	return i.Complete()
}

type InfectionManager struct {
	virus      *Virus
	infections chan *Infection
}

func NewInfectionManager(virus *Virus) *InfectionManager {
	infections := make(chan *Infection, 1e6)

	return &InfectionManager{
		virus:      virus,
		infections: infections,
	}
}

func (i *InfectionManager) CreateInfection(updateInfection UpdateInfection) {
	i.infections <- NewInfection(i.virus, updateInfection)
}

func (i *InfectionManager) PersistInfection(infection *Infection) {
	i.infections <- infection
}

func (i *InfectionManager) ProcessDay() {
	currentInfections := i.infections
	i.infections = make(chan *Infection, 1e6)

	for {
		select {
		case infection := <-currentInfections:
			infection.Progress()
		default:
			return
		}
	}
}
