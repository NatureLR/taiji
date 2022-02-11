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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
