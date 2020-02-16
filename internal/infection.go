package internal

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

type Infection struct {
	Name                 string
	MortalityProbability float64
	IncubationPeriod     int
	InfectiousIncubation bool
	ActivePeriod         int
	InfectiousActive     bool
	RecoveryPeriod       int
	InfectiousRecovery   bool
	UpdateInfection      UpdateInfection
	day                  int
}

func NewInfection(virus *Virus, updateInfection UpdateInfection) *Infection {
	// TODO(Sam): Slightly change values per infection here
	return &Infection{
		Name:                 virus.Name,
		MortalityProbability: virus.MortalityProbability,
		IncubationPeriod:     virus.IncubationPeriod,
		InfectiousIncubation: virus.InfectiousIncubation,
		ActivePeriod:         virus.ActivePeriod,
		InfectiousActive:     virus.InfectiousActive,
		RecoveryPeriod:       virus.RecoveryPeriod,
		InfectiousRecovery:   virus.InfectiousRecovery,
		UpdateInfection:      updateInfection,
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

	fmt.Println(random.Float64())
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
	virus            *Virus
	activeInfections chan *Infection
}

func NewInfectionManager(virus *Virus) *InfectionManager {
	activeInfections := make(chan *Infection, 1e6)

	return &InfectionManager{
		virus:            virus,
		activeInfections: activeInfections,
	}
}

func (i *InfectionManager) CreateInfection(updateInfection UpdateInfection) {
	i.activeInfections <- NewInfection(i.virus, updateInfection)
}

func (i *InfectionManager) ProcessDay() error {
	count := 0
	activeInfections := make(chan *Infection, 1e6)
	select {
	case infection := <-i.activeInfections:
		count++
		complete := infection.Progress()
		if !complete {
			activeInfections <- infection
		}
	default:
		break
	}

	i.activeInfections = activeInfections

	if count == 0 {
		return fmt.Errorf("Pandemic is over")
	}

	return nil
}
