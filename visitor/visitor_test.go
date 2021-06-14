package visitor

func ExampleVisitor() {
	c := &Computer{Parts: []ComputerPart{
		&Mouse{},
		&Monitor{},
		&Keyboard{},
	}}

	visitor := &ComputerVisitor{}

	//Output:
	//mouse
	//monitor
	//keyboard
	//computer
	c.Accept(visitor)
}
