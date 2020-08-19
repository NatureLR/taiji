package template

func init() {
	DefaultPool.add("readme", Readme, "README.md")
}

// Readme 文档模板
const Readme = `
# {{.project}}
`
