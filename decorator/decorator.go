package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Calc() int {
	return 0
}

type MultipleComponent struct {
	Val       int
	Component Component
}

func NewMultipleComponent(c Component, val int) Component {
	return &MultipleComponent{Component: c, Val: val}
}

func (m *MultipleComponent) Calc() int {
	return m.Val * m.Component.Calc()
}

type AddComponent struct {
	Val       int
	Component Component
}

func NewAddComponent(c Component, val int) Component {
	return &AddComponent{Component: c, Val: val}
}

func (a *AddComponent) Calc() int {
	return a.Val + a.Component.Calc()
}
