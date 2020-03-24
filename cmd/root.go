package cmd

import (
	"fmt"
	"os"

	tpl "github.com/NatureLingRan/go-project/pkg/template"
	"github.com/NatureLingRan/go-project/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-project",
	Short: "用创建go项目需要的文件",
	Long: fmt.Sprintf(`
	创建go项目需要的文件,如果指向创建某个文件执行: go-project init <文件类型>
	目前支持的有:%s
	`, tpl.AllKind()),
	Version: version.Print(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
