package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func ExecTemplate(text string, data any) []byte {
	buf := new(bytes.Buffer)

	tpl, err := template.New("").Parse(text)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if err = tpl.Execute(buf, data); err != nil {
		fmt.Println(err)
		return nil
	}

	return buf.Bytes()
}
