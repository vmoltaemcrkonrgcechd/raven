package back

import (
	"raven/pkg/converter"
	"raven/pkg/utils"
)

type Node struct {
	Name      string
	Type      string
	CanBeNil  bool
	Pkg       string
	Many      bool
	TableName string
	Tag       string
	Embedded  bool
	Children  []*Node
}

func NewNode(name, typ, pkg, tableName string, canBeNil, many bool, tag string, embedded bool) *Node {
	return &Node{
		Name:      name,
		Type:      typ,
		CanBeNil:  canBeNil,
		Pkg:       pkg,
		Many:      many,
		TableName: tableName,
		Tag:       tag,
		Embedded:  embedded,
	}
}

func (n *Node) AddChild(node *Node) *Node {
	n.Children = append(n.Children, node)
	return n
}

func (n *Node) Generate() []byte {
	return utils.ExecTemplate(StructTemplate, n)
}

const (
	QuantityName = "quantity"
	ContentName  = "content"
)

const (
	JSONTag = "json"
)

func (n *Node) PublicName() string {
	return converter.ToPascalCase(n.Name)
}

func (n *Node) PrivateName() string {
	return converter.ToCamelCase(n.Name)
}
