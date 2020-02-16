package internal

import (
	"fmt"
)

type Virus struct {
	Name                 string
	R0                   float64
	MortalityRate        float64
	IncubationPeriod     int
	InfectiousIncubation bool
	ActivePeriod         int
	InfectiousActive     bool
	RecoveryPeriod       int
	InfectiousRecovery   bool

	MortalityProbability    float64
	TransmissionProbability float64
	contagiousPeriod        int
}

func NewVirus(
	name string,
	r0 float64,
	mortalityRate float64,
	incubationPeriod int,
	infectiousIncubation bool,
	activePeriod int,
	infectiousActive bool,
	recoveryPeriod int,
	infectiousRecovery bool,
) *Virus {

	v := &Virus{
		Name:                 name,
		R0:                   r0,
		MortalityRate:        mortalityRate,
		IncubationPeriod:     incubationPeriod,
		InfectiousIncubation: infectiousIncubation,
		ActivePeriod:         activePeriod,
		InfectiousActive:     infectiousActive,
		RecoveryPeriod:       recoveryPeriod,
		InfectiousRecovery:   infectiousRecovery,
		contagiousPeriod:     0,
	}

	v.MortalityProbability = mortalityRate / float64(v.ActivePeriod)
	return v
}

func (v *Virus) ContagiousPeriod() int {
	if v.contagiousPeriod != 0 {
		return v.contagiousPeriod
	}

	contagiousPeriod := 0
	if v.InfectiousIncubation {
		contagiousPeriod += v.IncubationPeriod
	}

	if v.InfectiousActive {
		contagiousPeriod += v.ActivePeriod
	}

	if v.InfectiousRecovery {
		contagiousPeriod += v.RecoveryPeriod
	}

	v.contagiousPeriod = contagiousPeriod

	return v.contagiousPeriod
}

func (v *Virus) CalculateTransmissionProbability(averageContacts int) float64 {
	if v.TransmissionProbability != 0 {
		return v.TransmissionProbability
	}

	v.TransmissionProbability = v.R0 / float64(v.ContagiousPeriod()*averageContacts)
	return v.TransmissionProbability
}

func (v *Virus) Desc() string {
	return fmt.Sprintf(
		"Virus %q: R0 %.2f, Contagious Period %d days, Mortality Rate %.2f",
		v.Name,
		v.R0,
		v.ContagiousPeriod(),
		v.MortalityRate,
	)
}
