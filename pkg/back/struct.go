package back

import (
	"bytes"
	"fmt"
	"raven/pkg/value"
	"strings"
	"text/template"
)

type StructField struct {
	*value.Value
	Struct []*StructField
	Tag    string
}

func NewStructField(val *value.Value, tag string) *StructField {
	return &StructField{
		Value: val,
		Tag:   tag,
	}
}

const (
	JSONTag = "json"
)

func (s *StructField) Generate() []byte {
	buf := new(bytes.Buffer)

	template.Must(template.New("struct").Parse(Struct)).Execute(buf, s)

	return buf.Bytes()
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
