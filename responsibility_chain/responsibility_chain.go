package responsibility_chain

import "fmt"

type Manager interface {
	HaveRight(money float64) bool
	Audit()
}

type Chain struct {
	Manager   Manager
	Seccessor *Chain
}

func NewChain(manager Manager) *Chain {
	return &Chain{Manager: manager}
}

func (c *Chain) SetSeccessor(m Manager) {
	c.Seccessor = NewChain(m)
}

type ProjectManager struct{}
type DepManager struct{}
type GeneralManager struct{}

func (p *ProjectManager) HaveRight(money float64) bool {
	if money < 100 {
		return true
	}
	return false
}
func (p *DepManager) HaveRight(money float64) bool {
	if money < 500 {
		return true
	}
	return false
}
func (p *GeneralManager) HaveRight(money float64) bool {
	if money < 5000 {
		return true
	}
	return false
}

func (p *ProjectManager) Audit() {
	fmt.Println("project manager audit")
}
func (p *DepManager) Audit() {
	fmt.Println("dep manager audit")
}
func (p *GeneralManager) Audit() {
	fmt.Println("general manager audit")
}

func (c *Chain) Audit(money float64) {
	if c.Manager.HaveRight(money) {
		c.Manager.Audit()
	} else {
		if c.Seccessor == nil {
			fmt.Println("no body can audit")

		} else {
			c.Seccessor.Audit(money)
		}
	}
}
