package project

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/NatureLingRan/go-project/pkg/tools"
)

// Project 需要创建的对象
type Project struct {
	Name string
	Path string
}

// File 创建文件类型的方法
func (p *Project) File(filetype, content string) {
	fp := filepath.Join(p.Path, p.Name, filetype)
	if p.Path == "." {
		fp = filepath.Join(p.Path, filetype)
	}
	log.Println("创建文件:", fp)
	f, err := os.Create(fp)
	tools.CheckErr(err)
	_, err = f.WriteString(content)
	tools.CheckErr(err)
}

// Dir 创建文件夹
func (p *Project) Dir(dirname string) {
	fp := filepath.Join(p.Path, p.Name, dirname)
	if p.Path == "." {
		fp = dirname
	}
	log.Println("创建文件夹:", fp)
	tools.CheckErr(os.MkdirAll(fp, 0744))
}

// Parsecontent 创建模板文件的各个模板
func (p *Project) Parsecontent(content, project string) string {
	return strings.ReplaceAll(content, `{{.project}}`, project)
}

// New 创建对象
func New(name, path string) *Project {
	return &Project{
		Name: name,
		Path: path,
	}
}

// Create 创建一些文件和目录
func (p *Project) Create() {
	p.Dir("pkg")
	p.Dockerfile()
	p.VersionGo()
	p.K8s()
	p.Makefile()
	p.Gitignore()
	p.Readme()
	p.License()
	p.Corba()
}

// Update 更新创建的文件
func (p *Project) Update(args []string) {
	if len(args) == 0 {
		p.Create()
		return
	}
	arg := strings.ToLower(args[0])

	switch arg {
	case "makefile":
		p.Makefile()
	case "k8s":
		p.K8s()
	case "version":
		p.VersionGo()
	case "dockerfile":
		p.Dockerfile()
	case "license":
		p.License()
	case "readme":
		p.Readme()
	case "gitignore":
		p.Gitignore()
	default:
		fmt.Println("不支持的文件类型")
	}
}
