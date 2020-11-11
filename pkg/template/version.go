package template

func init() {
	Default.Add("version.go", Version, "cmd/version.go")
	Default.Add("versions.go", LibVersion, "pkg/versions/versions.go")
	Default.Add("description.go", Description, "pkg/versions/description.go")
}

// Version cmd/version.go模板
const Version = `
package cmd

import (
	"{{.importPath}}/pkg/versions"
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
	versionCmd.Flags().StringVarP(&versions.Format, "format", "f", "string", "版本信息输出的格式，目前有两种种:string,json")
	rootCmd.AddCommand(versionCmd)
}
`

// LibVersion cmd/version.go模板
const LibVersion = `
package versions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
)

// 通过go build -X 注入的参数
var (
	xVersion   string // Version 程序版本
	xGitCommit string // GitCommit git提交的hash
	xBuilt     string // Built 编译时间
)

// Info 版本信息
type Info struct {
	Version   string {{.backquoted}}json:"version"{{.backquoted}}
	GitCommit string {{.backquoted}}json:"gitCommit"{{.backquoted}}
	GoVersion string {{.backquoted}}json:"goVersion"{{.backquoted}}
	Built     string {{.backquoted}}json:"built"{{.backquoted}}
}

// New 返回*info对象
func New(version, gitcommit, built string) *Info {
	return &Info{
		Version:   version,
		GitCommit: gitcommit,
		GoVersion: runtime.Version(),
		Built:     built,
	}
}

// JSON 输出json
func (i *Info) JSON() string {
	ret, err := json.Marshal(i)
	if err != nil {
		log.Print("解析json错误")
		os.Exit(1)
	}
	return string(ret)
}

// Column 以一列的形式
func (i *Info) Column() string {
	return fmt.Sprintf("%s GitCommit:%s GoVersion:%s Built:%s", i.Version, i.GitCommit, i.GoVersion, i.Built)
}

// Strings 输出字符串
func (i *Info) Strings(format string) string {
	var s string
	switch format {
	case "json":
		s = i.JSON()
	case "column":
		s = i.Column()
	default:
		s = i.Column()
	}
	return s
}

func init() {
	Default = New(xVersion, xGitCommit, xBuilt)
}

var (
	// Default 输出的格式
	Default *Info
	// Format 输出的格式
	Format string
)

// Print 将版本输出到控制台
func Print() {
	fmt.Println(Default.Strings(Format))
}

// Strings 输出字符串
func Strings() string {
	return Default.Strings(Format)
}
`

// Description 版本描述
const Description = `
package versions

// ShortDescribe 简单的描述
var ShortDescribe = "用于创建GO项目的脚手架"

// LongDescribe 长描述
var LongDescribe = ""
// TODO 改为从README中动态加载
`
