package back

const (
	Struct = "{{.PublicName}} " +
		"{{if .Many}}[]{{end}}" +
		"{{if .CanBeNil}}*{{end}}" +
		"struct {\n" +
		"{{range .Children}}" +
		"{{if .Children}}" +
		"{{printf \"%s\" .Generate}}" +
		"{{else}}" +
		"{{if not .Embedded}}{{.PublicName}} {{end}}" +
		"{{if .Many}}[]{{end}}" +
		"{{if .CanBeNil}}*{{end}}" +
		"{{.Type}}" +
		"{{end}}" +
		"{{if ne .Tag \"\"}} `{{.Tag}}:\"{{.PrivateName}}\"` {{end}}\n" +
		"{{end}}}"
)
