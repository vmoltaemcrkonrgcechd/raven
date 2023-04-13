package back

type Module struct {
	Table string
}

func NewModule(table string) *Module {
	return &Module{
		Table: table,
	}
}
