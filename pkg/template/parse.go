package template

import (
	"strings"
)

// Parsecontent 创建模板文件的各个模板
func Parsecontent(content, project string) string {
	return strings.ReplaceAll(content, `{{.project}}`, project)
}
