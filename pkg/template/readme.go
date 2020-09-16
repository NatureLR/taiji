package template

func init() {
	Default.Add("readme", Readme, "README.md")
}

// Readme 文档模板
const Readme = `
# {{.project}}
`
