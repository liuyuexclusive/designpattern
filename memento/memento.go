package memento

type Memento interface{}

type Game struct {
	hp, mp int
}

type GameMemento struct {
	hp, mp int
}

func (g *Game) Play(a, b int) {
	g.hp += a
	g.mp += b
}

func (g *Game) Save() Memento {
	return &GameMemento{hp: g.hp, mp: g.mp}
}

func (g *Game) Load(me Memento) {
	m := me.(*GameMemento)
	g.hp = m.hp
	g.mp = m.mp
}
