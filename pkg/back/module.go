package back

import (
	"fmt"
	"raven/pkg/converter"
	"raven/pkg/postgres"
	"raven/pkg/utils"
	"strconv"
)

type Module struct {
	Table    string
	pg       *postgres.Postgres
	names    map[string]map[string]struct{}
	Entities []*Node
}

func NewModule(table string, pg *postgres.Postgres) *Module {
	return &Module{
		Table: table,
		pg:    pg,
		names: make(map[string]map[string]struct{}),
	}
}

func (mod *Module) Create(columns []string) error {
	return nil
}

func (mod *Module) Read(cmd ReadCommand) error {
	rootName := mod.checkName(EntityPkg, fmt.Sprintf("All%sDto", converter.ToPascalCase(mod.Table)), 0)

	root := NewNode(
		rootName, converter.StructType, EntityPkg, "", false, false,
		"", false)

	content := NewNode(
		ContentName, converter.StructType, "", mod.Table, false, true,
		JSONTag, false)

	root.AddChild(content).AddChild(NewNode(
		QuantityName, converter.IntType, "", "", false, false,
		JSONTag, false))

	_, err := mod.FillNode(content, cmd.Columns, cmd.Join)
	if err != nil {
		return err
	}

	mod.Entities = append(mod.Entities, root)

	return nil
}

func (mod *Module) Update(columns []string) error {
	return nil
}

func (mod *Module) Delete() error {
	return nil
}

func (mod *Module) FillNode(node *Node, columns []string, join []*Join) (*Node, error) {
	table, err := mod.pg.GetTable(node.TableName)
	if err != nil {
		return nil, err
	}

	var (
		column *postgres.Column
		typ    string
	)
	for _, name := range columns {
		if column, err = table.GetColumn(name); err != nil {
			return nil, err
		}

		if typ, err = converter.PgToGoType(column.Type); err != nil {
			return nil, err
		}

		node.AddChild(NewNode(
			name, typ, node.TableName, "", column.CanBeNil, false,
			JSONTag, false))
	}

	for _, i := range join {
		if len(i.Columns) > 0 {
			childNode := NewNode(
				i.Name(), converter.StructType, "", i.Table, false, i.Many,
				JSONTag, false)

			if childNode, err = mod.FillNode(childNode, i.Columns, i.Join); err != nil {
				return nil, err
			}

			node.AddChild(childNode)

			continue
		}

		if node, err = mod.FillNode(node, i.Columns, i.Join); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func (mod *Module) GenerateEntities() []byte {
	return utils.ExecTemplate(Entities, mod)
}

func (mod *Module) checkName(pkg, name string, n int) string {
	if _, ok := mod.names[pkg]; !ok {
		mod.names[pkg] = make(map[string]struct{})
		mod.names[pkg][name] = struct{}{}
		return name
	}

	newName := name
	if n != 0 {
		newName += strconv.Itoa(n)
	}

	if _, ok := mod.names[pkg][newName]; ok {
		return mod.checkName(pkg, name, n+1)
	}

	mod.names[pkg][newName] = struct{}{}

	return newName
}
