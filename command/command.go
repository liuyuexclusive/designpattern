package command

import "fmt"

type Command interface {
	Excute()
}

type MotherBoard struct {
}

func (m *MotherBoard) Start() {
	fmt.Println("start")
}

func (m *MotherBoard) Stop() {
	fmt.Println("stop")
}

type StartCommand struct {
	Board *MotherBoard
}

func (s *StartCommand) Excute() {
	s.Board.Start()
}

type StopCommand struct {
	Board *MotherBoard
}

func (s *StopCommand) Excute() {
	s.Board.Stop()
}

type Box struct {
	Button1, Button2 Command
}

func NewBox(c1, c2 Command) *Box {
	return &Box{c1, c2}
}
