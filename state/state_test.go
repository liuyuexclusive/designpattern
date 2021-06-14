package state

import "testing"

func TestState(t *testing.T) {
	c := &Context{Today: &Monday{}}

	(&Friday{}).Action((c))

	if c.Today.Val() != 5 {
		t.Errorf("got: %v,want %v", c.Today.Val(), 5)
	}
}
