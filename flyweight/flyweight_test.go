package flyweight

import "fmt"

func ExampleFlyweight() {

	//Output:
	//data from test.txt
	//data from test.txt
	factory := NewFactory()
	x1 := factory.Get("test.txt").Data

	fmt.Println(x1)

	x2 := factory.Get("test.txt").Data

	fmt.Println(x2)

}
