package visitor

import "fmt"

type ComputerPartVisitor interface {
	Visit(ComputerPart)
}

type ComputerPart interface {
	Accept(ComputerPartVisitor)
}

type Keyboard struct{}

type Mouse struct{}

type Monitor struct{}

type Computer struct {
	Parts []ComputerPart
}

func (k *Keyboard) Accept(c ComputerPartVisitor) {
	c.Visit(k)
}
func (k *Mouse) Accept(c ComputerPartVisitor) {
	c.Visit(k)
}
func (k *Monitor) Accept(c ComputerPartVisitor) {
	c.Visit(k)
}
func (k *Computer) Accept(c ComputerPartVisitor) {
	for _, v := range k.Parts {
		c.Visit(v)
	}
	c.Visit(k)
}

type ComputerVisitor struct {
}

func (c *ComputerVisitor) Visit(v ComputerPart) {
	switch v.(type) {
	case *Keyboard:
		fmt.Println("keyboard")
	case *Mouse:
		fmt.Println("mouse")
	case *Monitor:
		fmt.Println("monitor")
	case *Computer:
		fmt.Println("computer")
	}
}
