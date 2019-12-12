package create

import (
	tl "go-project/pkg/template"
	"go-project/pkg/tools"
	"log"
	"os"
	"path/filepath"
)

// Create 需要创建的对象
type Create struct {
	Name string
	Path string
}

// Dir 创建文件夹
func (c *Create) Dir(dirname string) {
	fp := filepath.Join(c.Path, c.Name, dirname)
	log.Println("创建文件夹:", fp)
	tools.CheckErr(os.MkdirAll(fp, 0744))
}

// Gitignore 创建.gitignore文件
func (c *Create) Gitignore() {
	c.cf(".gitignore", c.Name)
}

// Makefile  创建makefile
func (c *Create) Makefile() {
	c.cf("Dockerfile", tl.Parsecontent(tl.Makefile, c.Name))
}

// Dockerfile 创建dockerfile
func (c *Create) Dockerfile() {
	c.cf("Dockerfile", tl.Parsecontent(tl.Dockerfile, c.Name))
}

// K8s 创建k8s
func (c *Create) K8s() {
	c.cf("k8s.yaml", tl.Parsecontent(tl.K8s, c.Name))
}

func (c *Create) cf(filetype, content string) {
	fp := filepath.Join(c.Path, c.Name, filetype)
	log.Println("创建文件:", fp)
	f, err := os.Create(fp)
	tools.CheckErr(err)
	_, err = f.WriteString(content)
	tools.CheckErr(err)
}
