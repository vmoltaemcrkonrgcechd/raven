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

	SQLNamesTemplate = "{{range $i, $v := .}}{{if ne $i 0}},{{end}}\"{{$v.Name}}\"{{end}}"

	TypesTemplate = "{{range $i, $v := .}}{{if ne $i 0}},{{end}}{{$v.Type}}{{end}}"

	ConstructorTemplate = "return {{.Returns.Types}}{ {{.Parameters.Names}} }"

	CreateTemplate = "if {{.Returns.ErrNode.PrivateName}} = {{.Recipient.PrivateName}}.Sq.Insert(\"\\\"{{.TableName}}\\\"\").\n" +
		"{{$entity := index .Parameters 0}}" +
		"Columns({{$entity.Children.SQLNames}}).\n" +
		"Values({{range $i, $v := $entity.Children}}{{if ne $i 0}},{{end}}{{$entity.PrivateName}}.{{$v.PublicName}}{{end}}).\n" +
		"Suffix(\"RETURNING {{.Returns.IDNode.Name}}\")." +
		"QueryRow().Scan(&{{.Returns.IDNode.PrivateName}}); {{.Returns.ErrNode.PrivateName}} != nil {\n" +
		"return {{.Returns.Names}}\n}\n\n" +
		"return {{.Returns.Names}}"

	ConfigTemplate = "package config\n\n" +
		"type Config struct {\n" +
		"HTTPAddr   string `yaml:\"httpAddr\"`\nPgURL      string `env:\"PG_URL\"`\n}\n\n" +
		"func New() (*Config, error) {\n" +
		"cfg := new(Config)\n\n" +
		"err := cleanenv.ReadConfig(\"./config/config.yaml\", cfg)\n" +
		"if err != nil {\nreturn nil, err\n}\n\n" +
		"if err = godotenv.Load(); err != nil {\nreturn nil, err\n}\n\n" +
		"if err = cleanenv.ReadEnv(cfg); err != nil {\nreturn nil, err\n}\n\n" +
		"return cfg, nil\n}\n"

	ConfigYamlTemplate = "httpAddr: \":80\"\n"

	PostgresTemplate = "package postgres\n\n" +
		"type Postgres struct {\nSq squirrel.StatementBuilderType\n" +
		"DB *sql.DB\n}\n\n" +
		"func New(cfg *config.Config) (*Postgres, error) {\n" +
		"db, err := sql.Open(\"postgres\", cfg.PgURL)\n" +
		"if err != nil {\nreturn nil, err\n}\n\n" +
		"if err = db.Ping(); err != nil {\n" +
		"return nil, err\n}\n\n" +
		"return &Postgres{\nSq: squirrel.StatementBuilder.\n" +
		"PlaceholderFormat(squirrel.Dollar).RunWith(db),\n" +
		"DB: db,\n}, nil\n}\n"
)
