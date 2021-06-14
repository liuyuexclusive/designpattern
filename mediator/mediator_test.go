package mediator

func ExampleMediator() {
	cd := &CDDriver{}
	aCard := &AudioCard{}
	vCard := &VedioCard{}
	cpu := &CPU{}

	m := &Mediator{
		CPU:       cpu,
		AudioCard: aCard,
		VedioCard: vCard,
		CDDriver:  cd,
	}

	//Output:
	//play music by audio card
	//play vedio by vedio card
	cd.Read("music,vedio", m)
}
