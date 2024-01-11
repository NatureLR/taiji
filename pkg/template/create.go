package template

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"text/template"

	"github.com/naturelr/taiji/pkg/tools"
	"github.com/naturelr/taiji/pkg/versions"
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

	goversion := runtime.Version()
	versionRegex := regexp.MustCompile(`go(\d+\.\d+)`)
	matches := versionRegex.FindStringSubmatch(goversion)
	if len(matches) >= 2 {
		goversion = matches[1]
	} else {
		fmt.Println("Failed to extract version.")
	}

	err = tmpl.Execute(f, map[string]string{
		"project":          peoject,
		"importPath":       impotPath,
		"backquoted":       "`",
		"ShortDescribe":    versions.ShortDescribe,
		"LongDescribe":     versions.LongDescribe,
		"GoVersion":        goversion,
		"LeftDoubleBrace":  "{{", // {{
		"RightDoubleBrace": "}}", // }}
	})
	tools.CheckErr(err)
}
