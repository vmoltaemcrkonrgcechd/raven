package back

import (
	"raven/pkg/utils"
	"raven/pkg/value"
)

type Node struct {
	*value.Value
	Tag      string
	Embedded bool
	Children []*Node
}

func NewNode(val *value.Value, tag string, embedded bool) *Node {
	return &Node{
		Value:    val,
		Tag:      tag,
		Embedded: embedded,
	}
}

func (n *Node) AddChild(node *Node) *Node {
	n.Children = append(n.Children, node)
	return n
}

func (n *Node) Generate() []byte {
	return utils.ExecTemplate(Struct, n)
}

const (
	QuantityName = "quantity"
	ContentName  = "content"
)

const (
	JSONTag = "json"
)
