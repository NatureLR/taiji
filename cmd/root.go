package cmd

import (
	"fmt"
	"os"

	"github.com/NatureLingRan/go-project/version"
	"github.com/spf13/cobra"
)

var (
	projectPath string
	projectName string
)

var rootCmd = &cobra.Command{
	Use:     "go-project",
	Short:   "用创建go项目需要的文件",
	Long:    `用于创建go项目的一些文件如Dockerfile,目录结构等`,
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

func init() {

	rootCmd.PersistentFlags().StringVar(&projectName, "name", "", "项目名字")
	rootCmd.PersistentFlags().StringVar(&projectPath, "path", ".", "项目创建的路径")

}
