package project

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/NatureLingRan/go-project/pkg/template"
	"github.com/NatureLingRan/go-project/pkg/tools"
	"github.com/spf13/cobra"
)

// Init  如果没有指定创建文件就创建所有文件，否则就创建指定的文件
func Init(cmd *cobra.Command, args []string) {

	project := filepath.Base(os.Getenv("PWD"))

	if len(args) == 0 {
		for n, t := range template.GetDefaul() {
			c := parsecontent(t.Content, project)
			Create(t.Path, n, c)
		}
		return
	}

	t := template.GetDefaul()[args[0]]
	content := parsecontent(t.Content, t.Path)
	Create(t.Path, args[0], content)
}

// Create 创建
func Create(path, name, content string) {
	if path != "." {
		log.Println("创建文件夹:", path)
		tools.CheckErr(os.MkdirAll(path, 0744))
	}
	log.Println("创建文件:", name)
	f, err := os.Create(filepath.Join(path, name))
	tools.CheckErr(err)
	_, err = f.WriteString(content)
	tools.CheckErr(err)
}

func importPath() string {
	path := strings.ReplaceAll(os.Getenv("PWD"), os.Getenv("GOPATH")+"/src/", "")
	return strings.Replace(path, "\\", "/", -1) //将\替换成/
}

func parsecontent(content, project string) string {
	c := strings.ReplaceAll(content, `{{.project}}`, project)
	c = strings.ReplaceAll(c, `{{.importPath}}`, importPath())
	c = strings.ReplaceAll(c, `{{.importPath}}`, importPath())
	return c
}
