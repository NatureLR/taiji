package template

func init() {
	Default.Add("version.go", Version, "cmd/version.go")
}

// Version version.go模板
const Version = `
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// 通过go build -X 注入的参数
var (
	// Version 程序版本
	version string
	// GitCommit git提交的hash
	gitCommit string
	// GoVersion 编译的go版本
	goVersion string
	// Built 编译时间
	built string
)

// Print 打印版本
func versionString() string {
	if versionFormat == "json" {
		v, err := json.Marshal(map[string]string{
			"version":    version,
			"gitCommmit": gitCommit,
			"goVersion":  goVersion,
			"built":      built,
		})
		if err != nil {
			log.Print("解析json错误")
			os.Exit(1)
		}
		return string(v)
	}
	return fmt.Sprintf("%s GitCommit:%s GoVersion:%s Built:%s", version, gitCommit, goVersion, built)
}

func versionPrint() {
	fmt.Println(versionString())
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "打印出版本号",
	Run: func(cmd *cobra.Command, args []string) {
		versionPrint()
	},
}

var versionFormat string

func init() {
	versionCmd.Flags().StringVarP(&versionFormat, "format", "f", "string", "版本信息输出的格式，目前有两种种:string,json")
	rootCmd.AddCommand(versionCmd)
}
`
