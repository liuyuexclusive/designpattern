package command

func ExampleCommand() {
	board := &MotherBoard{}
	start := &StartCommand{board}
	stop := &StopCommand{board}

	//Output:
	//start
	//stop
	//stop
	//start

	box1 := NewBox(start, stop)

	box1.Button1.Excute()
	box1.Button2.Excute()

	box2 := NewBox(stop, start)

	box2.Button1.Excute()
	box2.Button2.Excute()

}
