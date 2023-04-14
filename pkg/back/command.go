package back

import "encoding/json"

type CommandJSON struct {
	Type  string          `json:"type"`
	Table string          `json:"table"`
	Info  json.RawMessage `json:"info"`
}

type CreateOrUpdateCommand struct {
	Columns []string `json:"columns"`
}

type ReadCommand struct {
	Columns []string `json:"columns"`
}

const (
	CreateType = "create"
	ReadType   = "read"
	UpdateType = "update"
	DeleteType = "delete"
)
