package back

import (
	"os"
	"os/exec"
)

type Back struct {
	Dir     string
	GoMod   string
	Modules map[string]*Module
}

func New(dir, goMod string) *Back {
	return &Back{
		Dir:     dir,
		GoMod:   goMod,
		Modules: make(map[string]*Module),
	}
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
