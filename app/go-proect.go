package app

import (
	"go-project/pkg/create"
	"go-project/pkg/tools"
	"log"
)

type project interface {
	AddDir([]string)
	Create()
}

func creteProject(p project) {
	p.AddDir([]string{"app", "pkg"})
	p.Create()

}

// GoProject 创建project
func GoProject(name, path string) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(tools.Trace("%v", e.(error).Error()))

		}
	}()
	p := &create.Create{
		Name: name,
		Path: path,
	}
	creteProject(p)
}
