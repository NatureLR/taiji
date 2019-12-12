package app

import (
	"go-project/pkg/create"
	"go-project/pkg/tools"
	"log"
)

type project interface {
	addDir([]string)
	dirs() string
	files() string
	careteFile()
	careteDir()
}

// CreateProject 创建project
func CreateProject(name, path string) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(tools.Trace("%v", e.(error).Error()))

		}
	}()

	p := &create.Create{
		Name: name,
		Path: path,
	}

	p.Dir("app")
	p.Dir("pkg")
	p.Gitignore()
	p.Makefile()
	p.Dockerfile()
}
