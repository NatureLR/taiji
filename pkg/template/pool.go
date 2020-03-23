package template

// Pool 模板对象集合
type Pool struct {
	Templates map[string]*Template
}

// Template  模板对象
type Template struct {
	Content string
	Path    string
}

func (p *Pool) init() {
	p.Templates = make(map[string]*Template)
}

func (p *Pool) add(name, tpl, path string) {
	if p.Templates == nil {
		p.init()
	}
	if path == "" {
		path = "."
	}
	p.Templates[name] = &Template{
		Content: tpl,
		Path:    path,
	}
}

// Get 获取模板池的模板对象
func (p *Pool) Get(name string) *Template {
	return p.Templates[name]
}

// GetPath 获取模板池的模板对象
func (p *Pool) GetPath(name string) string {
	return p.Templates[name].Content
}

// GetContent 获取模板池的模板对象
func (p *Pool) GetContent(name string) *Template {
	return p.Templates[name]
}

// GetDefaul 获取默认所有的
func GetDefaul() map[string]*Template {
	return DefaultPool.Templates
}

// DefaultPool 默认的模板池子，所有的模板都在里面
var DefaultPool Pool
