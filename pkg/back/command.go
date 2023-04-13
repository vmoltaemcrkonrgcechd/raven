package back

import "encoding/json"

type CommandJSON struct {
	Type  string          `json:"type"`
	Table string          `json:"table"`
	Info  json.RawMessage `json:"info"`
}

type AddOrEditCommand struct {
	Columns []string `json:"columns"`
}
