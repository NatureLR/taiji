package project

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	tl "github.com/NatureLingRan/go-project/pkg/template"
	"github.com/NatureLingRan/go-project/pkg/tools"
)

// Project 需要创建的对象
type Project struct {
	Name string
	Path string
}

// Gitignore 创建.gitignore文件
func (p *Project) Gitignore() {
	p.File(".gitignore", p.Name)
}

// Makefile  创建makefile
func (p *Project) Makefile() {
	p.File("Makefile", p.Parsecontent(tl.Makefile, p.Name))
}

// Dockerfile 创建dockerfile
func (p *Project) Dockerfile() {
	p.File("Dockerfile", p.Parsecontent(tl.Dockerfile, p.Name))
}

// K8s 创建k8s
func (p *Project) K8s() {
	p.File("k8s.yaml", p.Parsecontent(tl.K8s, p.Name))
}

// VersionGo 创建version.go文件
func (p *Project) VersionGo() {
	p.Dir("version")
	p.File("version/version.go", p.Parsecontent(tl.Version, p.Name))
}

// Readme 创建Readme.md文件
func (p *Project) Readme() {
	p.File("README.md", p.Parsecontent(tl.Readme, p.Name))
}

// License 创建License文件
func (p *Project) License() {
	p.File("LICENSE", tl.License)
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
		fp = p.Path
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
