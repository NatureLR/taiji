package template

func init() {
	DefaultPool.add("README.md", Readme, "")
}

// Readme 文档模板
const Readme = `
# {{.project}}
`
