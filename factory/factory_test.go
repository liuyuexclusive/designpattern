package factory

import "fmt"

func ExampleFactory() {
	//Output:
	//5
	//5
	var f1 OperatorFactory = &OperatorPlusFactory{}
	o1 := f1.Create()
	o1.SetA(3)
	o1.SetB(2)
	fmt.Println(o1.Result())
	var f2 OperatorFactory = &OperatorMinusFactory{}
	o2 := f2.Create()
	o2.SetA(10)
	o2.SetB(5)
	fmt.Println(o2.Result())
}
