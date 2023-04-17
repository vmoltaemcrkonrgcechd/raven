package back

type Repo struct {
	Node        *Node
	Methods     []*RepoMethod
	Constructor *Method
}

func NewRepo(node *Node) *Repo {
	return &Repo{
		Node: node,
		Constructor: NewMethod(
			"New"+node.PublicName(), nil, node.Children, Nodes{node}, ConstructorTemplate,
		),
	}
}

func (r *Repo) AddMethod(method *RepoMethod) *Repo {
	r.Methods = append(r.Methods, method)
	return r
}
