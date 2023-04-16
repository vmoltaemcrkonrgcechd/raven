package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func ExecTemplate(text string, data any) string {
	buf := new(bytes.Buffer)

	tpl, err := template.New("").Parse(text)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if err = tpl.Execute(buf, data); err != nil {
		fmt.Println(err)
		return ""
	}

	return buf.String()
}
