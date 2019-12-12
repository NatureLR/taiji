package cmd

import (
	"go-project/app"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化go项目",
	Run: func(cmd *cobra.Command, args []string) {
		app.CreateProject(projectName, projectPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVar(&projectPath, "path", ".", "项目创建的路径")
	initCmd.PersistentFlags().StringVar(&projectName, "name", "", "项目名字")
}