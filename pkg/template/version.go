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
	rootCmd.AddCommand(versionCmd)
}
`

// LibVersion cmd/version.go模板
const LibVersion = `
package versions

import (
	"fmt"
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

func (i *Info) strings() string {
	return fmt.Sprintf("%s GitCommit:%s GoVersion:%s Built:%s", i.Version, i.GitCommit, i.GoVersion, i.Built)
}

func Strings() string {
	return Default.strings()
}

func Print() {
	fmt.Println(Default.strings())
}

func init() {
	Default = New(xVersion, xGitCommit, xBuilt)
}

var Default *Info
`

// Description 版本描述
const Description = `
package versions

// ShortDescribe 简单的描述
var ShortDescribe = "用于创建GO项目的脚手架"

// LongDescribe 长描述
var LongDescribe = ""
`
