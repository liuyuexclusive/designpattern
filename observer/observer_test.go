package observer

func ExampleObserver() {
	s := NewSubject()
	s.AddObserver(&ObserverA{})
	s.AddObserver(&ObserverA{})
	s.AddObserver(&ObserverB{})
	s.UpdateContext("hello")

	//Output:
	// observer a catch new context: hello
	// observer a catch new context: hello
	// observer b catch new context: hello
}
