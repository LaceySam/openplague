package internal

type Group struct {
	General    map[string]*Person
	Contagious map[string]*Person
	Sick       map[string]*Person

	AverageEncounters     int
	AverageSickEncounters int
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
