package cmd

import (
	"github.com/NatureLingRan/go-project/pkg/project"
	"github.com/spf13/cobra"
)

var projectName string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化go项目",
	Run:   project.Init,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
