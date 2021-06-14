package interpreter_type

import (
	"log"
	"strconv"
	"strings"
)

type Node interface {
	Calc() int
}

type ValNode struct {
	Val int
}

func (v *ValNode) Calc() int {
	return v.Val
}

type MinusNode struct {
	Left, Right Node
}

func (m *MinusNode) Calc() int {
	return m.Left.Calc() - m.Right.Calc()
}

type PlusNode struct {
	Left, Right Node
}

func (p *PlusNode) Calc() int {
	return p.Left.Calc() + p.Right.Calc()
}

type Parser struct {
	Exp   []string
	Index int
	Pre   Node
}

func NewParser(command string) *Parser {
	res := &Parser{}
	res.Exp = strings.Split(command, " ")
	return res
}

func (p *Parser) Parse() {
	for {
		if p.Index >= len(p.Exp) {
			break
		}
		c := p.Exp[p.Index]
		switch c {
		case "+":
			p.Index++
			c = p.Exp[p.Index]
			p.Pre = &PlusNode{Left: p.Pre, Right: NewValNode(c)}
		case "-":
			p.Index++
			c = p.Exp[p.Index]
			p.Pre = &MinusNode{Left: p.Pre, Right: NewValNode(c)}
		default:
			p.Pre = NewValNode(c)
		}
		p.Index++
	}
}

func NewValNode(c string) Node {
	v, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}
	return &ValNode{v}
}
