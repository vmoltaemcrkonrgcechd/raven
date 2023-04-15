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

type Join struct {
	Type    string   `json:"type"`
	Table   string   `json:"table"`
	Columns []string `json:"columns"`
	Use     string   `json:"use"`
	As      string   `json:"as"`
	Many    bool     `json:"many"`
	Join    []*Join  `json:"join"`
}

func (j *Join) Name() string {
	if j.As != "" {
		return j.As
	}

	return j.Table
}

type ReadCommand struct {
	Columns []string `json:"columns"`
	Join    []*Join  `json:"join"`
}

const (
	CreateType = "create"
	ReadType   = "read"
	UpdateType = "update"
	DeleteType = "delete"
)
