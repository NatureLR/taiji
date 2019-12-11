package app

import (
	"fmt"
	"go-project/pkg/create"
	"go-project/pkg/tools"
)

// CreateProject 创建p project。。。
func CreateProject() {
	defer func() {
		if e := recover(); e != nil {
			log := tools.Trace("%v", e.(error).Error())
			fmt.Println(log)
		}
	}()

	p := &create.Create{
		Name: "test",
		Path: "testpath",
	}

	p.Dir("app")
	p.Dir("pkg")
	p.Gitignore()
	p.Makefile()
	p.Dockerfile()
}
