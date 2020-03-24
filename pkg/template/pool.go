package template

import (
	"strings"
)

// Pool 模板对象集合
type Pool struct{ Templates map[string]*Template }

func (p *Pool) add(name, tpl, path string) {
	if p.Templates == nil {
		p.init()
	}

	if path == "" {
		path = "."
	}

	p.Templates[name] = &Template{
		content: tpl,
		path:    path,
	}
}

func (p *Pool) init() { p.Templates = make(map[string]*Template) }

// Template  模板对象
type Template struct {
	content string
	path    string
}

// Path 返回改模板需要创建的路径
func (t *Template) Path() string { return t.path }

// Content 返回模板内容
func (t *Template) Content() string { return t.content }

// DefaultPool 默认的模板池子，所有的模板都在里面
var DefaultPool Pool

// GetDefaul 获取默认所有的
func GetDefaul() map[string]*Template { return DefaultPool.Templates }

// DefaulKind 获取默认所有的
func DefaulKind(kind string) *Template { return DefaultPool.Templates[kind] }

// AllKind 所有的某班类型
func AllKind() string {
	var kinds []string

	for kind := range DefaultPool.Templates {
		kinds = append(kinds, kind)
	}
	return strings.Join(kinds, ",")
}
