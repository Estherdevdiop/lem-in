package lemin

type Room struct {
	name     string
	adjacent []string
	visit    bool
}

type Ant struct {
	Num      int
	Position int
	Arrive   bool
	Movable  bool
	Route    []string
}

type Parcour struct {
	Route   []string
	Ants    []*Ant
	SetMove int
}

func NewRoom(name string) *Room {
	return &Room{
		name:  name,
		visit: false,
	}
}

func NewParcour(r []string) *Parcour {
	return &Parcour{
		Route:   r,
		SetMove: 1,
	}
}

func NewAnt(num int) *Ant {
	return &Ant{
		Num:      num,
		Position: 1,
		Arrive:   false,
		Movable:  false,
		Route:    []string{},
	}
}

func InitParcours() {
	for _, route := range result {
		p := NewParcour(route)
		Parcours = append(Parcours, p)
	}
}

func InitAnt() {
	for i := 1; i <= NbAnt; i++ {
		ant := NewAnt(i)
		AllAnts = append(AllAnts, ant)
	}
}