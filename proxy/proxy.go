package proxy

import "fmt"

type Subject interface {
	do()
}

type RealSubject struct{}

func (r *RealSubject) do() {
	fmt.Println("action")
}

type ProxySubject struct {
	real *RealSubject
}

func NewProxy() *ProxySubject {
	return &ProxySubject{real: &RealSubject{}}
}

func (p *ProxySubject) do() {
	fmt.Println("begin")
	p.real.do()
	fmt.Println("end")
}
