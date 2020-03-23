package template

func init() {
	DefaultPool.add(".gitignore", Gitignore, "")
}

// Gitignore git的屏蔽模板
const Gitignore = `{{.project}}
.vscode
`
