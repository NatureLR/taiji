package create

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	tl "github.com/NatureLingRan/go-project/pkg/template"
	"github.com/NatureLingRan/go-project/pkg/tools"
)

// Create 需要创建的对象
type Create struct {
	Name string
	Path string
	Dirs []string
}

// Gitignore 创建.gitignore文件
func (c *Create) Gitignore() {
	c.File(".gitignore", c.Name)
}

// Makefile  创建makefile
func (c *Create) Makefile() {
	c.File("Makefile", c.Parsecontent(tl.Makefile, c.Name))
}

// Dockerfile 创建dockerfile
func (c *Create) Dockerfile() {
	c.File("Dockerfile", c.Parsecontent(tl.Dockerfile, c.Name))
}

// K8s 创建k8s
func (c *Create) K8s() {
	c.File("k8s.yaml", c.Parsecontent(tl.K8s, c.Name))
}

// VersionGo 创建version.go文件
func (c *Create) VersionGo() {
	c.Dir("version")
	c.File("version/version.go", c.Parsecontent(tl.Version, c.Name))
}

// File 创建文件类型的方法
func (c *Create) File(filetype, content string) {
	fp := filepath.Join(c.Path, c.Name, filetype)
	log.Println("创建文件:", fp)
	f, err := os.Create(fp)
	tools.CheckErr(err)
	_, err = f.WriteString(content)
	tools.CheckErr(err)
}

// Dir 创建文件夹
func (c *Create) Dir(dirname string) {
	fp := filepath.Join(c.Path, c.Name, dirname)
	log.Println("创建文件夹:", fp)
	tools.CheckErr(os.MkdirAll(fp, 0744))
}

// Parsecontent 创建模板文件的各个模板
func (c *Create) Parsecontent(content, project string) string {
	return strings.ReplaceAll(content, `{{.project}}`, project)
}

// AddDir 需要创建的dir
func (c *Create) AddDir(dir []string) {
	c.Dirs = append(c.Dirs, dir...)
}

// Create 创建一些文件和目录
func (c *Create) Create() {
	for _, d := range c.Dirs {
		c.Dir(d)
	}
	c.Dockerfile()
	c.K8s()
	c.Makefile()
	c.Gitignore()
	c.VersionGo()
}
