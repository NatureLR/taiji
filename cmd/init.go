package cmd

import (
	"go-project/app"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化go项目",
	Run: func(cmd *cobra.Command, args []string) {
		app.Create()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
