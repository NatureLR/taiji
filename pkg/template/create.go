package template

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/NatureLingRan/go-project/pkg/tools"
	"github.com/NatureLingRan/go-project/pkg/versions"
)

// CreateTpl 创建模板的接口
type CreateTpl interface {
	Path() string
	Content() string
}

// Create 解析模板创建文件
func Create(c CreateTpl, mod string) {
	peoject := filepath.Base(os.Getenv("PWD"))
	impotPath := tools.ImportPath(mod)

	if c.Path() != "" {
		dir := filepath.Dir(c.Path())
		tools.CheckErr(os.MkdirAll(dir, 0744))
	}
	log.Println("创建:", c.Path())
	f, err := os.Create(filepath.Join(c.Path()))
	tools.CheckErr(err)

	tmpl, err := template.New("goProject").Parse(c.Content())
	tools.CheckErr(err)

	err = tmpl.Execute(f, map[string]string{
		"project":       peoject,
		"importPath":    impotPath,
		"backquoted":    "`",
		"ShortDescribe": versions.ShortDescribe,
		"LongDescribe":  versions.LongDescribe,
	})
	tools.CheckErr(err)
}
