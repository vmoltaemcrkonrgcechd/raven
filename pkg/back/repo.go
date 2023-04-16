package back

type Repo struct {
	Node *Node
}

func NewRepo(node *Node) *Repo {
	return &Repo{node}
}
