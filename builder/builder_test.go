package builder

import "fmt"

func Example_Builder() {
	//Output:
	// 55.5
	// name: veg burger
	// price: 25.5
	// pack: wrapper
	// name: coke
	// price: 30.0
	// pack: bottle
	// 85.5
	// name: chicken burger
	// price: 50.5
	// pack: wrapper
	// name: pepsi
	// price: 35.0
	// pack: bottle
	mb := &MealBuilder{}
	m1 := mb.VegBurgerCoke()
	fmt.Println(m1.Price())
	m1.Print()

	m2 := mb.ChickenBurgerPepsi()
	fmt.Println(m2.Price())
	m2.Print()
}
