package template

import (
	"errors"
	"strings"
)

// Pool 模板对象集合
type Pool struct {
	// 文件类型：模板对象
	Templates map[string]*Template
}

// Add kind文件类型 path为路径和文件名的组合
func (p *Pool) Add(kind, tpl, path string) {
	if p.Templates == nil {
		p.init()
	}
	p.Templates[kind] = &Template{
		content: tpl,
		path:    path,
	}
}

func (p *Pool) init() { p.Templates = make(map[string]*Template) }

// Get 获取指定kind的模板
func (p *Pool) Get(kind string) *Template {
	if _, ok := p.Templates[kind]; !ok {
		panic(errors.New("不支持的类型"))
	}
	return p.Templates[kind]
}

// All 获取所有模板对象
func (p *Pool) All() map[string]*Template { return p.Templates }

// Allkind 获取所有kind
func (p *Pool) Allkind() string {
	var kinds []string
	for kind := range p.Templates {
		kinds = append(kinds, kind)
	}
	return strings.Join(kinds, ",")
}

// Template  模板对象
type Template struct {
	// 模板内容
	content string
	// 创建的路径包括文件名
	path string
}

// Path 返回改模板需要创建的路径
func (t *Template) Path() string { return t.path }

// Content 返回模板内容
func (t *Template) Content() string { return t.content }

// Default 默认的模板池，所有的默认模板都在里面
var Default Pool
