package facade

import "fmt"

func ExampleFacade() {
	api := NewAPI(&AModuleAPIImpl{}, &BModuleAPIImpl{})

	res := api.Test()

	//Output:
	//a module running,b module running
	fmt.Println(res)
}
