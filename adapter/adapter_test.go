package adapter

import "fmt"

func ExampleAdapter() {
	target := NewAdapter(&AdapteeImpl{})

	res := target.Request()

	//Output:
	//specific request
	fmt.Println(res)
}
