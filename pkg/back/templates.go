package back

const (
	Struct = "{{.PublicName}} {{if .Many}}[]{{end}}{{if .CanBeNil}}*{{end}}struct {\n" +
		"{{range .Struct}}" +
		"{{if .Struct}}{{printf \"%s\" .Generate}}{{$length := len .Tag}}{{if ne $length 0}} `{{.Tag}}:\"{{.PublicName}}\"`\n{{end}}" +
		"{{else}}{{.PublicName}} {{if .Many}}[]{{end}}{{if .CanBeNil}}*{{end}}{{.Type}}" +
		"{{$length := len .Tag}}{{if ne $length 0}} `{{.Tag}}:\"{{.PublicName}}\"`{{end}}\n" +
		"{{end}}{{end}}}"
)
