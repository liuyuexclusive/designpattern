package simple_factory

func ExampleSimple() {
	t1 := NewAPI(0)
	t2 := NewAPI(1)
	//Output:
	// hi james
	// hello james
	t1.Say("james")
	t2.Say("james")
}
