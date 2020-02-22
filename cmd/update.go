package cmd

import (
	"os"
	"path/filepath"

	"github.com/NatureLingRan/go-project/pkg/project"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "用于升级文件",
	Long: `用于升级Makefile，dockerfile等文件
		   执行 go-project update 将升级所有的文件
		   如: go-project makefile
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pn := filepath.Base(os.Getenv("PWD"))
		project.New(pn, projectPath).Update(args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringVar(&projectPath, "path", ".", "项目创建的路径")
}
