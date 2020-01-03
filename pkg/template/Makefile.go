package template

// Makefile 模板
const Makefile = `
build: 
	go build -ldflags "-s -w" .
docker:
	docker build -t {{.project}}:latest -f Dockerfile .
`
