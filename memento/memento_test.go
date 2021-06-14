package memento

import "testing"

func TestMemento(t *testing.T) {
	g := &Game{hp: 100, mp: 100}
	m := g.Save()
	g.Play(-20, -80)
	if g.hp != 80 || g.mp != 20 {
		t.Errorf("wrong")
	}
	g.Load(m)
	if g.hp != 100 || g.mp != 100 {
		t.Errorf("wrong")
	}

}
