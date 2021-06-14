package composite

func ExampleComposite() {
	//Output:
	//4
	//	2
	//		1
	//		3
	//	6
	//		5
	//		6
	root := NewTreeNode(4)
	l := NewTreeNode(2)
	r := NewTreeNode(6)
	l.AddChild(NewTreeNode(1))
	l.AddChild(NewTreeNode(3))
	r.AddChild(NewTreeNode(5))
	r.AddChild(NewTreeNode(6))
	root.AddChild(l)
	root.AddChild(r)
	root.Print("")
}
