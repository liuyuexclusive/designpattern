package responsibility_chain

func ExampleChain() {
	//Output:
	// project manager audit
	// dep manager audit
	// general manager audit
	// no body can audit
	m1 := &ProjectManager{}
	m2 := &DepManager{}
	m3 := &GeneralManager{}

	c1 := NewChain(m1)
	c1.SetSeccessor(m2)
	c1.Seccessor.SetSeccessor(m3)

	c1.Audit(50)
	c1.Audit(300)
	c1.Audit(700)
	c1.Audit(10000)
}
