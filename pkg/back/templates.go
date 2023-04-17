package back

const (
	StructTemplate = "{{.PublicName}} " +
		"{{if .Many}}[]{{end}}" +
		"{{if .CanBeNil}}*{{end}}" +
		"struct {\n" +
		"{{range .Children}}" +
		"{{if .Children}}" +
		"{{.Generate}}" +
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
		"\ntype {{.Generate}}\n{{end}}"

	RepoTemplate = "\ntype {{.Node.Generate}}\n" +
		"\n{{.Constructor.Generate}}\n" +
		"{{range .Methods}}" +
		"\n{{.Generate}}\n" +
		"{{end}}"

	MethodTemplate = "func {{if .Recipient}}" +
		"({{.Recipient.PrivateName}} {{.Recipient.PublicName}}) {{end}}" +
		"{{.Name}}({{.Parameters.ToParameters}})" +
		"({{.Returns.ToParameters}}) {\n" +
		"{{.GenerateBody}}" +
		"\n}"

	ParameterTemplate = "{{.PrivateName}} {{if .Many}}[]{{end}}" +
		"{{if .CanBeNil}}*{{end}}" +
		"{{if and (ne .Pkg \"\") (ne .Pkg .IgnorePkg)}}{{.Pkg}}.{{end}}" +
		"{{.Type}}"

	ParametersTemplate = "{{range $i, $v := .}}{{if ne $i 0}},{{end}}{{$v.ToParameter}}{{end}}"

	NamesTemplate = "{{range $i, $v := .}}{{if ne $i 0}},{{end}}{{$v.PrivateName}}{{end}}"

	TypesTemplate = "{{range $i, $v := .}}{{if ne $i 0}},{{end}}{{$v.Type}}{{end}}"

	ConstructorTemplate = "return {{.Returns.Types}}{ {{.Parameters.Names}} }"
)
