package decorator

import "fmt"

func ExampleDecorator() {
	a := &ConcreteComponent{}
	b := NewAddComponent(a, 10)
	res := NewMultipleComponent(b, 8)

	//Output:
	//80
	fmt.Println(res.Calc())
}
