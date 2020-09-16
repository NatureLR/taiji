package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/NatureLingRan/go-project/pkg/project"
	tpl "github.com/NatureLingRan/go-project/pkg/template"
	"github.com/spf13/cobra"
)

func initfFunc(cmd *cobra.Command, args []string) {
	defer func() {
		if e := recover(); e != nil {
			log.Fatal(e)
		}
	}()

	// 没有指定创建的文件类型就创建所有,制定了就创建指定的
	if len(args) == 0 {
		for _, t := range tpl.Default.All() {
			project.Create(t, mod)
		}
	} else {
		for _, kind := range args {
			kind = strings.ToLower(kind)
			t := tpl.Default.Get(kind)
			project.Create(t, mod)
		}
	}
}

var initLong = fmt.Sprintf(`
	如果只创建某个文件执行: go-project init <文件类型>
	目前支持文件类型有:%s

	例子：
	创建一个项目:
	以项目的名字创建一个文件夹，然后在文件夹里执行go-project init 

	在一个已经存在的项目中仅仅只是想创建个dockerfile:
	在你想创建的位置执行 go-project init Dockerfil <其他文件>
	`, tpl.Default.Allkind())

var initCmd = &cobra.Command{
	Use:   "init <类型>",
	Short: "初始化go项目",
	Long:  initLong,
	Run:   initfFunc,
}

var mod string

func init() {
	initCmd.Flags().StringVarP(&mod, "mod", "m", "", "如果不在GOPATH中创建项目需要指定mod name")
	rootCmd.AddCommand(initCmd)
}
