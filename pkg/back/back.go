package back

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"raven/pkg/postgres"
)

type Back struct {
	PgURL   string
	Dir     string
	GoMod   string
	Modules map[string]*Module
	pg      *postgres.Postgres
}

func New(dir, goMod, pgURL string, pg *postgres.Postgres) *Back {
	return &Back{
		Dir:     dir,
		GoMod:   goMod,
		PgURL:   pgURL,
		Modules: make(map[string]*Module),
		pg:      pg,
	}
}

type Config struct {
	PgURL    string        `json:"pgURL"`
	GoMod    string        `json:"goMod"`
	DirName  string        `json:"dirName"`
	Commands []CommandJSON `json:"commands"`
}

func ReadConfig(name string) (*Back, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	var pg *postgres.Postgres
	if pg, err = postgres.New(cfg.PgURL); err != nil {
		return nil, err
	}

	back := New(cfg.DirName, cfg.GoMod, cfg.PgURL, pg)

	return back, back.ExecCommands(cfg.Commands)
}

const (
	Perm = 0600
)

const (
	MainPkg     = "main"
	RepoPkg     = "repo"
	RoutesPkg   = "routes"
	EntityPkg   = "entity"
	AppPkg      = "app"
	PostgresPkg = "postgres"
	ConfigPkg   = "config"
)

const (
	MainPath     = "/cmd/app"
	RepoPath     = "/internal/" + RepoPkg
	RoutesPath   = "/internal/" + RoutesPkg
	EntityPath   = "/internal/" + EntityPkg
	AppPath      = "/internal/" + AppPkg
	PostgresPath = "/pkg/" + PostgresPkg
	ConfigPath   = "/" + ConfigPkg
)

func (b *Back) ProjectInit() error {
	err := os.MkdirAll(b.Dir+MainPath, Perm)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(b.Dir+RepoPath, Perm); err != nil {
		return nil
	}

	if err = os.MkdirAll(b.Dir+RoutesPath, Perm); err != nil {
		return nil
	}

	if err = os.MkdirAll(b.Dir+EntityPath, Perm); err != nil {
		return nil
	}

	if err = os.MkdirAll(b.Dir+AppPath, Perm); err != nil {
		return nil
	}

	if err = os.MkdirAll(b.Dir+PostgresPath, Perm); err != nil {
		return nil
	}

	if err = os.MkdirAll(b.Dir+ConfigPath, Perm); err != nil {
		return nil
	}

	if err = b.GoModInit(); err != nil {
		return err
	}

	if err = b.InstallDependencies(); err != nil {
		return err
	}

	return nil
}

func (b *Back) GoModInit() error {
	cmd := exec.Command("go", "mod", "init", b.GoMod)
	cmd.Dir = b.Dir
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (b *Back) InstallDependencies() error {
	var (
		err   error
		cmd   *exec.Cmd
		names = []string{
			"github.com/gofiber/fiber/v2",
			"github.com/gofiber/swagger",
			"github.com/lib/pq",
			"github.com/Masterminds/squirrel",
			"github.com/ilyakaznacheev/cleanenv",
		}
	)

	for _, name := range names {
		cmd = exec.Command("go", "get", name)

		cmd.Dir = b.Dir

		if err = cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Back) ExecCommand(cmdJSON CommandJSON) error {
	switch cmdJSON.Type {
	case CreateType:
		var cmd CreateOrUpdateCommand
		if err := json.Unmarshal(cmdJSON.Info, &cmd); err != nil {
			return err
		}
		return b.Create(cmdJSON.Table, cmd.Columns)

	case ReadType:
		var cmd ReadCommand
		if err := json.Unmarshal(cmdJSON.Info, &cmd); err != nil {
			return err
		}
		return b.Read(cmdJSON.Table, cmd)

	case UpdateType:
		var cmd CreateOrUpdateCommand
		if err := json.Unmarshal(cmdJSON.Info, &cmd); err != nil {
			return err
		}
		return b.Update(cmdJSON.Table, cmd.Columns)

	case DeleteType:
		return b.Delete(cmdJSON.Table)

	default:
		return errors.New("unknown command: " + cmdJSON.Type)
	}
}

func (b *Back) ExecCommands(commands []CommandJSON) error {
	var err error

	for _, cmd := range commands {
		if err = b.ExecCommand(cmd); err != nil {
			return err
		}
	}

	return nil
}

func (b *Back) GetModule(table string) *Module {
	if _, ok := b.Modules[table]; !ok {
		b.Modules[table] = NewModule(table, b.pg)
	}

	return b.Modules[table]
}

func (b *Back) Create(table string, columns []string) error {
	return b.GetModule(table).Create(columns)
}

func (b *Back) Read(table string, cmd ReadCommand) error {
	return b.GetModule(table).Read(cmd)
}

func (b *Back) Update(table string, columns []string) error {
	return b.GetModule(table).Update(columns)
}

func (b *Back) Delete(table string) error {
	return b.GetModule(table).Delete()
}

func (b *Back) Exec() error {
	err := b.ProjectInit()
	if err != nil {
		return err
	}

	for k, v := range b.Modules {
		if err = os.WriteFile(b.filenameForEntity(k),
			append([]byte("package "+EntityPkg+"\n"), v.GenerateEntities()...),
			Perm); err != nil {
			return err
		}

		if err = os.WriteFile(b.filenameForRepo(k),
			append([]byte("package "+RepoPkg+"\n"), v.GenerateRepo()...),
			Perm); err != nil {
			return err
		}
	}

	return nil
}

func (b *Back) filenameForEntity(name string) string {
	return fmt.Sprintf("%s%s/%s.go", b.Dir, EntityPath, name)
}

func (b *Back) filenameForRepo(name string) string {
	return fmt.Sprintf("%s%s/%s.go", b.Dir, RepoPath, name)
}
