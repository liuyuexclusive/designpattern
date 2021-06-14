package mediator

import (
	"fmt"
	"strings"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) Read(data string, m *Mediator) {
	c.Data = data
	m.Change(c)
}

type AudioCard struct {
	// Data string
}

func (a *AudioCard) Play(data string) {
	fmt.Printf("play %s by audio card\n", data)
}

type VedioCard struct {
	// Data string
}

func (a *VedioCard) Play(data string) {
	fmt.Printf("play %s by vedio card\n", data)
}

type CPU struct {
	Audio string
	Vedio string
}

func (c *CPU) Process(data string, m *Mediator) {
	slice := strings.Split(data, ",")
	c.Audio = slice[0]
	c.Vedio = slice[1]
	m.Change(c)
}

type Mediator struct {
	CPU       *CPU
	AudioCard *AudioCard
	VedioCard *VedioCard
	CDDriver  *CDDriver
}

func (m *Mediator) Change(i interface{}) {
	switch s := i.(type) {
	case *CDDriver:
		m.CPU.Process(s.Data, m)
	case *CPU:
		m.AudioCard.Play(s.Audio)
		m.VedioCard.Play(s.Vedio)
	}
}
