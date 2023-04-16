package back

type Repo struct {
	Node    *Node
	Methods []*Method
}

func NewRepo(node *Node) *Repo {
	return &Repo{
		Node: node,
	}
}

func (r *Repo) AddMethod(method *Method) *Repo {
	r.Methods = append(r.Methods, method)
	return r
}
