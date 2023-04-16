package back

const (
	StructTemplate = "{{.PublicName}} " +
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
		"{{if ne .Pkg \"\"}}{{.Pkg}}.{{end}}" +
		"{{.Type}}" +
		"{{end}}" +
		"{{if ne .Tag \"\"}} `{{.Tag}}:\"{{.PrivateName}}\"` {{end}}\n" +
		"{{end}}}"

	EntitiesTemplate = "{{range .Entities}}" +
		"\ntype {{printf \"%s\" .Generate}}\n{{end}}"

	RepoTemplate = "\ntype {{printf \"%s\" .Node.Generate}}\n"
)
