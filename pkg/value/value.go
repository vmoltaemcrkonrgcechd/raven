package value

type Value struct {
	Name      string
	Type      string
	CanBeNil  bool
	Pkg       string
	Many      bool
	TableName string
}

func New(name, typ, pkg, tableName string, canBeNil, many bool) *Value {
	return &Value{
		Name:      name,
		Type:      typ,
		CanBeNil:  canBeNil,
		Pkg:       pkg,
		Many:      many,
		TableName: tableName,
	}
}
