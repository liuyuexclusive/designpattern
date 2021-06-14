package interpreter_type

import "fmt"

func ExampleInterpreter() {
	parser := NewParser("1 + 2 + 3 - 4 + 5 - 6")
	parser.Parse()
	//Output:
	//1
	fmt.Println(parser.Pre.Calc())
}
