package app

import (
	"fmt"
	"go-project/pkg/create"
	"go-project/pkg/tools"
)

type project interface {
	dirs() string
	files() string
	careteFile()
	careteDir()
}

// CreateProject 创建project
func CreateProject(name, path string) {
	defer func() {
		if e := recover(); e != nil {
			log := tools.Trace("%v", e.(error).Error())
			fmt.Println(log)
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
