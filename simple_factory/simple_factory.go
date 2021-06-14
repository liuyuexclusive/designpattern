package simple_factory

import "fmt"

type API interface {
	Say(name string)
}

type Hi struct{}

func (h *Hi) Say(name string) {
	fmt.Printf("hi %s\n", name)
}

type Hello struct{}

func (h *Hello) Say(name string) {
	fmt.Printf("hello %s\n", name)
}

func NewAPI(t int) API {
	if t == 0 {
		return &Hi{}
	} else {
		return &Hello{}
	}
}
