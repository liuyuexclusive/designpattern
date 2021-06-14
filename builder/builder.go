package builder

import "fmt"

type Packing interface {
	Pack() string
}

type Item interface {
	Name() string
	Price() float64
	Pack() Packing
}

type Wrapper struct{}

func (w *Wrapper) Pack() string {
	return "wrapper"
}

type Bottle struct{}

func (b *Bottle) Pack() string {
	return "bottle"
}

type Burger struct{}

func (b *Burger) Pack() Packing {
	return &Wrapper{}
}

type ColdDrink struct{}

func (b *ColdDrink) Pack() Packing {
	return &Bottle{}
}

type VegBurger struct {
	*Burger
}

func (v *VegBurger) Price() float64 {
	return 25.5
}

func (v *VegBurger) Name() string {
	return "veg burger"
}

type ChickenBurger struct {
	*Burger
}

func (v *ChickenBurger) Price() float64 {
	return 50.5
}

func (v *ChickenBurger) Name() string {
	return "chicken burger"
}

type Coke struct {
	*ColdDrink
}

func (c *Coke) Price() float64 {
	return 30.0
}

func (c *Coke) Name() string {
	return "coke"
}

type Pepsi struct {
	*ColdDrink
}

func (c *Pepsi) Price() float64 {
	return 35.0
}

func (c *Pepsi) Name() string {
	return "pepsi"
}

type Meal struct {
	Items []Item
}

func (m *Meal) AddItem(item Item) {
	m.Items = append(m.Items, item)
}

func (m *Meal) Price() float64 {
	var res float64

	for _, v := range m.Items {
		res += v.Price()
	}
	return res
}

func (m *Meal) Print() {
	for _, v := range m.Items {
		fmt.Printf("name: %s\n", v.Name())
		fmt.Printf("price: %.1f\n", v.Price())
		fmt.Printf("pack: %s\n", v.Pack().Pack())
	}
}

type MealBuilder struct{}

func (m *MealBuilder) VegBurgerCoke() *Meal {
	res := &Meal{}
	res.AddItem(&VegBurger{})
	res.AddItem(&Coke{})
	return res
}

func (m *MealBuilder) ChickenBurgerPepsi() *Meal {
	res := &Meal{}
	res.AddItem(&ChickenBurger{})
	res.AddItem(&Pepsi{})
	return res
}
