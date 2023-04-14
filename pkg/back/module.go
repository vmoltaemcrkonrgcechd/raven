package back

type Module struct {
	Table string
}

func NewModule(table string) *Module {
	return &Module{
		Table: table,
	}
}

func (mod *Module) Create(columns []string) error {
	return nil
}

func (mod *Module) Read(columns []string) error {
	return nil
}

func (mod *Module) Update(columns []string) error {
	return nil
}

func (mod *Module) Delete() error {
	return nil
}
