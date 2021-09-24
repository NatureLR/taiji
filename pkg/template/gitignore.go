package template

func init() {
	Default.Add("gitignore", Gitignore, ".gitignore")
}

// Gitignore git的屏蔽模板
const Gitignore = `{{.project}}
.vscode
bin/
build/
.DS_Store
`
