package back

import (
	"os"
	"os/exec"
)

type Back struct {
	Dir   string
	GoMod string
}

func New(dir, goMod string) *Back {
	return &Back{
		Dir:   dir,
		GoMod: goMod,
	}
}

const (
	Perm = 0600
)

const (
	MainPkg     = "main"
	UseCasePkg  = "usecase"
	RepoPkg     = "repo"
	HTTPPkg     = "http"
	EntityPkg   = "entity"
	AppPkg      = "app"
	PostgresPkg = "postgres"
	ConfigPkg   = "config"
)

const (
	MainPath     = "/cmd/app"
	UseCasePath  = "/internal/" + UseCasePkg
	RepoPath     = "/internal/usecase/" + RepoPkg
	HTTPPath     = "/internal/controller/" + HTTPPkg
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

	if err = os.MkdirAll(b.Dir+HTTPPath, Perm); err != nil {
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

	return nil
}

func (b *Back) GoModInit() error {
	cmd := exec.Command("go", "mod", "init", b.GoMod)
	cmd.Dir = "./" + b.Dir
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
