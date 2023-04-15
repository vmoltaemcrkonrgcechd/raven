package back

import (
	"fmt"
	"raven/pkg/value"
	"strings"
)

type StructField struct {
	*value.Value
	Struct []*StructField
}

func (s *StructField) Log(n int) {
	var space = strings.Repeat("  ", n)
	msg := fmt.Sprint(space, s.TableName, ".", s.Name)

	if s.Many {
		msg += "[]"
	}

	fmt.Println(msg)

	for _, str := range s.Struct {
		str.Log(n + 1)
	}
}

func NewStructField(val *value.Value) *StructField {
	return &StructField{
		Value: val,
	}
}
