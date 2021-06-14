package composite

import "fmt"

type Node interface {
	GetVal() int
	GetNdoes() []Node
	AddChild(Node)
	Print(string)
}

type TreeNode struct {
	Val   int
	Nodes []*TreeNode
}

func (t *TreeNode) GetVal() int {
	return t.Val
}

func (t *TreeNode) GetNdoes() []Node {
	var res []Node
	for _, v := range t.Nodes {
		res = append(res, v)
	}
	return res
}

func (t *TreeNode) AddChild(n Node) {
	t.Nodes = append(t.Nodes, n.(*TreeNode))
}

func (t *TreeNode) Print(prefix string) {
	if t == nil {
		return
	}
	fmt.Printf("%s%d\n", prefix, t.GetVal())
	prefix += "\t"
	for _, v := range t.GetNdoes() {
		v.Print(prefix)
	}
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
