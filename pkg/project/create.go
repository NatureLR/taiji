package project

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/NatureLingRan/go-project/pkg/tools"
)

func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		def := filepath.Join(home, "go")
		if filepath.Clean(def) == filepath.Clean(runtime.GOROOT()) {
			// Don't set the default GOPATH to GOROOT,
			// as that will trigger warnings from the go tool.
			return ""
		}
		return def
	}
	return ""
}

// 判断是否在GOPATH中
// 如果是就使用GOPATH的路径
// 不是就需要指定mod
func importPath(path string) string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	goPath := defaultGOPATH()

	if strings.HasPrefix(pwd, goPath) {
		path = strings.ReplaceAll(pwd, goPath+"/src/", "")
	}
	if path == "" {
		log.Fatal("当前不在在GOPATH中,使用--mod或者-m 指定mod名字")
	}
	return strings.Replace(path, "\\", "/", -1) //将\替换成/
}

// CreateTpl 创建模板的接口
type CreateTpl interface {
	Path() string
	Content() string
}

// Create 解析模板创建文件
func Create(c CreateTpl, mod string) {
	peoject := filepath.Base(os.Getenv("PWD"))
	impotPath := importPath(mod)

	if c.Path() != "" {
		dir := filepath.Dir(c.Path())
		tools.CheckErr(os.MkdirAll(dir, 0744))
	}
	log.Println("创建:", c.Path())
	f, err := os.Create(filepath.Join(c.Path()))
	tools.CheckErr(err)

	tmpl, err := template.New("goProject").Parse(c.Content())
	tools.CheckErr(err)

	tmpl.Execute(f, map[string]string{
		"project":    peoject,
		"importPath": impotPath,
	})
}
