package back

import (
	"errors"
	"raven/pkg/postgres"
	"raven/pkg/utils"
	"raven/pkg/value"
)

type Module struct {
	Table string
	pg    *postgres.Postgres
}

func NewModule(table string, pg *postgres.Postgres) *Module {
	return &Module{
		Table: table,
		pg:    pg,
	}
}

func (mod *Module) Create(columns []string) error {
	return nil
}

func (mod *Module) Read(cmd ReadCommand) error {
	str, err := mod.joinStruct(mod.Table, mod.Table, false, false, cmd.Columns, cmd.Join)
	if err != nil {
		return err
	}

	str.Log(0)

	return nil
}

func (mod *Module) Update(columns []string) error {
	return nil
}

func (mod *Module) Delete() error {
	return nil
}

func (mod *Module) newStruct(tableName, structName string, canBeNil, many bool, columns []string) (*StructField, error) {
	table, err := mod.pg.GetTable(tableName)
	if err != nil {
		return nil, err
	}

	structField := NewStructField(value.New(structName, structName, EntityPkg, "", canBeNil, many))

	for _, columnName := range columns {
		var column *postgres.Column
		if column, err = table.GetColumn(columnName); err != nil {
			return nil, err
		}

		var typ string
		if typ, err = utils.PgToGoType(column.Type); err != nil {
			return nil, err
		}

		structField.Struct = append(structField.Struct,
			NewStructField(value.New(column.Name, typ, "", tableName, column.CanBeNil, false)))
	}

	return structField, nil
}

func (mod *Module) joinStruct(tableName, structName string,
	canBeNil, many bool, columns []string, join []*Join) (*StructField, error) {

	var (
		structField *StructField
		err         error
	)
	structField, err = mod.newStruct(tableName, structName, canBeNil, many, columns)
	if err != nil {
		return nil, err
	}

	queue := new([]*Join)

	*queue = append(*queue, join...)

	for _, j := range join {
		if j.Use != "" {
			var table *postgres.Table
			if table, err = mod.pg.GetTable(tableName); err != nil {
				return nil, err
			}

			var column *postgres.Column
			if column, err = table.GetColumn(j.Use); err != nil {
				return nil, err
			}

			var child *StructField
			child, err = mod.joinStruct(j.Table, j.Name(), column.CanBeNil, j.Many, j.Columns, j.Join)
			if err != nil {
				return nil, err
			}

			if len(j.Columns) == 0 {
				structField.Struct = append(structField.Struct, child.Struct...)
			} else {
				structField.Struct = append(structField.Struct, child)
			}

			continue
		}

		return nil, errors.New("вам нужно указать, какое поле использовать для объединения таблиц")
	}

	return structField, nil
}
