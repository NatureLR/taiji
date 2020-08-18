package cmd

import (
	"fmt"
	"os"

	tpl "github.com/NatureLingRan/go-project/pkg/template"
	"github.com/NatureLingRan/go-project/version"
	"github.com/spf13/cobra"
)

const short string = "用于创建GO项目的脚手架"

var long string = fmt.Sprintf(`
	创建go项目需要的一些文件
	如果只创建某个文件执行: go-project init <文件类型>
	目前支持文件类型有:%s

	例子：
	创建一个项目:
	以项目的名字创建一个文件夹，然后在文件夹里执行go-project init 

	在一个已经存在的项目中仅仅只是想创建个dockerfile:
	在你想创建的位置执行 go-project init Dockerfil <其他文件>
	`, tpl.AllKind())

var rootCmd = &cobra.Command{
	Use:     "go-project",
	Short:   short,
	Long:    long,
	Version: version.Print(),
}

// Execute 将所有的子命令加入到根命令并设置适当的flag
// 这是main.main()调用的,只调用一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
