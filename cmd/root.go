package cmd

import (
	"fmt"
	"os"

	"github.com/naturelr/taiji/pkg/versions"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "taiji",
	Short:   versions.ShortDescribe,
	Version: versions.Strings(),
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
	rootCmd.PersistentFlags().StringVar(&versions.Format, "format", "string", "版本信息输出的格式，目前有两种:string,json")
}
