package internal

type Group struct {
	General    map[string]*Person
	Contagious map[string]*Person
	Sick       map[string]*Person

	// When not sick, how many encounters do you have
	AverageEncounters int
	// When Sick can you still encounter people in this group
	SickEncounters bool
}

func NewGroup(sickEncounters bool) *Group {
	return &Group{
		General:        map[string]*Person{},
		Contagious:     map[string]*Person{},
		Sick:           map[string]*Person{},
		SickEncounters: sickEncounters,
	}
}

func (g *Group) AddPerson(person *Person) {
	g.General[person.Name] = person
}

func (g *Group) Event(event Event) {
	switch event.Kind {
	case CONTAGIOUS:
		g.Contagious[event.Name] = g.General[event.Name]

	case SICK:
		g.Sick[event.Name] = g.General[event.Name]
		delete(g.General, event.Name)
		delete(g.Contagious, event.Name)

	case RECOVERED:
		g.General[event.Name] = g.Sick[event.Name]
		delete(g.Sick, event.Name)

	case DEATH:
		delete(g.General, event.Name)
		delete(g.Contagious, event.Name)
		delete(g.Sick, event.Name)
	}
}

func (g *Group) Encounters() int {
	encounters := random.Intn(len(g.General))

	// Set a maximium
	if encounters > 100 {
		encounters = 100
	}

	return encounters
}

func (g *Group) Encounter(sick *Person) {
	// Random number of people from pool
	encounters := g.Encounters()
	encountered := 0

	for encountered < encounters {
		pick := random.Intn(len(g.General))
		i := 0
		for _, target := range g.General {
			i++
			if i != pick {
				continue
			}

			if random.Float64() < sick.Infection.TransmissionProbability {
				target.Infect()
			}

		}

		encountered++
	}
}

func (g *Group) ProcessSickEncounters() {
	for _, sick := range g.Contagious {
		g.Encounter(sick)
	}

	if !g.SickEncounters {
		return
	}

	for _, sick := range g.Sick {
		g.Encounter(sick)
	}
}

type City struct {
	Population   *Group
	HouseHolds   []*Group
	CommuteLines []*Group
	Companies    []*Group
}

func NewCity(householdCount int, commuteLineCount int, companyCount int) *City {
	city := &City{Population: NewGroup(false)}

	households := make([]*Group, householdCount)
	for i := 0; i < householdCount; i++ {
		households[i] = NewGroup(true)
	}

	city.HouseHolds = households

	commuteLines := make([]*Group, commuteLineCount)
	for i := 0; i < commuteLineCount; i++ {
		commuteLines[i] = NewGroup(false)
	}

	city.CommuteLines = commuteLines

	companies := make([]*Group, companyCount)
	for i := 0; i < companyCount; i++ {
		companies[i] = NewGroup(false)
	}

	city.Companies = companies

	return city
}

// Add person to city, random household and if they work, random commute line and random company
func (c *City) AddPerson(person *Person) {
	c.Population.AddPerson(person)

	// Add to random household
	c.HouseHolds[random.Intn(len(c.HouseHolds))].AddPerson(person)

	// Assume 80% of population work or in school
	if random.Float64() < 0.8 {
		c.CommuteLines[random.Intn(len(c.CommuteLines))].AddPerson(person)
		c.Companies[random.Intn(len(c.Companies))].AddPerson(person)
	}
}
