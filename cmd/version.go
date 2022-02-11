package cmd

import (
	"github.com/naturelr/taiji/pkg/versions"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "打印出版本号",
	Run: func(cmd *cobra.Command, args []string) {
		versions.Print()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
