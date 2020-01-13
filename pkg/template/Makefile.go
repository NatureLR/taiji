package template

// Makefile 模板
const Makefile = `
path = $$(go list)

version = -X '$(path)/version.Version=$$(git tag --sort=taggerdate |tail -1)'

goversion = -X '$(path)/version.GoVersion=$$(go version | awk '{printf($$3)}')'

gitcommit = -X '$(path)/version.GitCommit=$$(git rev-parse HEAD)'

built = -X '$(path)/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'

ldflag = "-s -w $(version) $(goversion) $(gitcommit) $(built)"

#ldflag = "-s -w -X rancher-restart/version.Version=V1.1.1 -X 'rancher-restart/version.GoVersion=$$(go version | awk '{printf($$3)}')' -X 'rancher-restart/version.GitCommit=$$(git rev-parse HEAD)' -X 'rancher-restart/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'"

build: 

	go build -ldflags $(ldflag) -x .

docker:

	docker build -t {{.project}}:latest -f Dockerfile .

`
