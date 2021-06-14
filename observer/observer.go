package observer

import "fmt"

type Observer interface {
	Update(*Subject)
}

type Subject struct {
	Observers []Observer
	Context   string
}

type ObserverA struct{}

type ObserverB struct{}

func (o *ObserverA) Update(s *Subject) {
	fmt.Printf("observer a catch new context: %s\n", s.Context)
}

func (o *ObserverB) Update(s *Subject) {
	fmt.Printf("observer b catch new context: %s\n", s.Context)
}

func (s *Subject) AddObserver(o Observer) {
	s.Observers = append(s.Observers, o)
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) Notify() {
	for _, v := range s.Observers {
		v.Update(s)
	}
}

func (s *Subject) UpdateContext(ctx string) {
	s.Context = ctx
	s.Notify()
}
