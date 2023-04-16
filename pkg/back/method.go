package back

import "raven/pkg/utils"

type Method struct {
	Name       string
	Recipient  *Node
	Parameters Nodes
	Returns    Nodes
}

func NewMethod(name string, recipient *Node, parameters Nodes, returns Nodes) *Method {
	return &Method{
		Name:       name,
		Recipient:  recipient,
		Parameters: parameters,
		Returns:    returns,
	}
}

func (m *Method) Generate() string {
	return utils.ExecTemplate(MethodTemplate, m)
}
