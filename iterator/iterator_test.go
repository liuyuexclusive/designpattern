package iterator

import "fmt"

func ExampleIterator() {
	var c Collection = []string{"a", "b", "c"}
	it := c.Iter()
	//Output:
	//a
	//b
	//c
	//a
	//b
	//c
	for it.CanNext() {
		fmt.Println(it.Next())
	}

	it2 := c.Iter()

	for it2.CanNext() {
		fmt.Println(it2.Next())
	}
}
