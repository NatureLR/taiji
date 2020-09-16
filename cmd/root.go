package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-project",
	Short:   "用于创建GO项目的脚手架",
	Version: versionString(),
}

// Execute 将所有的子命令加入到根命令并设置适当的flag
// 这是main.main()调用的,只调用一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&versionFormat, "format", "string", "版本信息输出的格式，目前有两种种:string,json")
}
