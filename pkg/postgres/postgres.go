package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Tables map[string]*Table
}

func New(pgURL string) (*Postgres, error) {
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if rows, err = db.Query(
		"SELECT c.table_name, c.column_name, c.data_type, c.is_nullable, LENGTH(COALESCE(c.column_default, '')) > 0, " +
			"COALESCE(tc.constraint_type, ''), COALESCE(kcu2.table_name, ''), COALESCE(kcu2.column_name, '') " +
			"FROM information_schema.columns c " +
			"LEFT JOIN information_schema.key_column_usage kcu ON c.table_name = kcu.table_name AND c.column_name = kcu.column_name " +
			"LEFT JOIN information_schema.table_constraints tc USING (constraint_name) " +
			"LEFT JOIN information_schema.referential_constraints USING (constraint_name) " +
			"LEFT JOIN information_schema.key_column_usage kcu2 ON unique_constraint_name = kcu2.constraint_name " +
			"WHERE c.table_schema = 'public'"); err != nil {
		return nil, err
	}

	var (
		tableName string
		canBeNil  string
		table     *Table
		pg        = &Postgres{Tables: make(map[string]*Table)}
	)

	for rows.Next() {
		column := new(Column)

		if err = rows.Scan(
			&tableName,
			&column.Name,
			&column.Type,
			&canBeNil,
			&column.HasDefault,
			&column.ConstraintType,
			&column.RefTable,
			&column.RefColumn,
		); err != nil {
			return nil, err
		}

		if canBeNil == "YES" {
			column.CanBeNil = true
		}

		table = pg.getTable(tableName)

		table.addColumn(column)
	}

	return pg, nil
}

func (pg *Postgres) getTable(name string) *Table {
	if _, ok := pg.Tables[name]; !ok {
		pg.Tables[name] = &Table{}
	}

	return pg.Tables[name]
}

type Table struct {
	Columns []*Column
}

func (t *Table) addColumn(column *Column) {
	t.Columns = append(t.Columns, column)
}

type Column struct {
	Name           string
	Type           string
	CanBeNil       bool
	HasDefault     bool
	ConstraintType string
	RefTable       string
	RefColumn      string
}
