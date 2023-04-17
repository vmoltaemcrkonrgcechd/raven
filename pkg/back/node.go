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
	Children  Nodes
	IgnorePkg string
}

func NewNode(name, typ, pkg, tableName string, canBeNil, many bool, tag string, embedded bool, ignorePkg string) *Node {
	return &Node{
		Name:      name,
		Type:      typ,
		CanBeNil:  canBeNil,
		Pkg:       pkg,
		Many:      many,
		TableName: tableName,
		Tag:       tag,
		Embedded:  embedded,
		IgnorePkg: ignorePkg,
	}
}

func (n *Node) AddChild(node *Node) *Node {
	n.Children = append(n.Children, node)
	return n
}

func (n *Node) Generate() string {
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

func (n *Node) ToParameter() string {
	return utils.ExecTemplate(ParameterTemplate, n)
}

type Nodes []*Node

func (n Nodes) ToParameters() string {
	return utils.ExecTemplate(ParametersTemplate, n)
}

func (n Nodes) Names() string {
	return utils.ExecTemplate(NamesTemplate, n)
}

func (n Nodes) Types() string {
	return utils.ExecTemplate(TypesTemplate, n)
}

func (n Nodes) ErrNode() *Node {
	for _, node := range n {
		if node == ErrNode {
			return node
		}
	}

	return nil
}

func (n Nodes) IDNode() *Node {
	for _, node := range n {
		if node == IDNode {
			return node
		}
	}

	return nil
}
