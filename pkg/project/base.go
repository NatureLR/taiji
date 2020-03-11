package project

import (
	"os"
	"strings"

	tl "github.com/NatureLingRan/go-project/pkg/template"
)

// Gitignore 创建.gitignore文件
func (p *Project) Gitignore() {
	p.File(".gitignore", p.Parsecontent(tl.Gitignore, p.Name))
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

// Corba 创建main.go和cmd文件夹下的root.go
func (p *Project) Corba() {
	p.Dir("cmd")

	rootGo := strings.ReplaceAll(p.Parsecontent(tl.RootGo, p.Name), `{{.importPath}}`, importPath())
	p.File("/cmd/root.go", rootGo)

	mainGo := strings.ReplaceAll(p.Parsecontent(tl.MainGo, p.Name), `{{.importPath}}`, importPath())
	p.File("main.go", mainGo)
}

func importPath() string {

	path := strings.ReplaceAll(os.Getenv("PWD"), os.Getenv("GOPATH")+"/src/", "")

	return strings.Replace(path, "\\", "/", -1) //将\替换成/
}
