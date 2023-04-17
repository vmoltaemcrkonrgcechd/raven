package back

import (
	"raven/pkg/utils"
)

type Method struct {
	Name       string
	Recipient  *Node
	Parameters Nodes
	Returns    Nodes
	Body       string
}

func NewMethod(name string, recipient *Node, parameters Nodes, returns Nodes, body string) *Method {
	return &Method{
		Name:       name,
		Recipient:  recipient,
		Parameters: parameters,
		Returns:    returns,
		Body:       body,
	}
}

func (m *Method) Generate() string {
	return utils.ExecTemplate(MethodTemplate, m)
}

func (m *Method) GenerateBody() string {
	return utils.ExecTemplate(m.Body, m)
}

type RepoMethod struct {
	*Method
	TableName string
}

func NewRepoMethod(method *Method, tableName string) *RepoMethod {
	return &RepoMethod{
		Method:    method,
		TableName: tableName,
	}
}

func (m *RepoMethod) Generate() string {
	return utils.ExecTemplate(MethodTemplate, m)
}

func (m *RepoMethod) GenerateBody() string {
	return utils.ExecTemplate(m.Body, m)
}
