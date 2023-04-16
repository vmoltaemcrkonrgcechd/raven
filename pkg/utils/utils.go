package utils

import (
	"bytes"
	"text/template"
)

func ExecTemplate(text string, data any) []byte {
	buf := new(bytes.Buffer)

	tpl, err := template.New("").Parse(text)
	if err != nil {
		return nil
	}

	if err = tpl.Execute(buf, data); err != nil {
		return nil
	}

	return buf.Bytes()
}
