package template

// Makefile 模板
const Makefile = `
project = $$(go list)

version = -X '$(project)/version.Version=$$(git tag --sort=taggerdate |tail -1)'

goversion = -X '$(project)/version.GoVersion=$$(go version | awk '{printf($$3)}')'

gitcommit = -X '$(project)/version.GitCommit=$$(git rev-parse HEAD)'

built = -X '$(project)/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'

ldflag = "-s -w $(version) $(goversion) $(gitcommit) $(built)"

#ldflag = "-s -w -X rancher-restart/version.Version=V1.1.1 -X 'rancher-restart/version.GoVersion=$$(go version | awk '{printf($$3)}')' -X 'rancher-restart/version.GitCommit=$$(git rev-parse HEAD)' -X 'rancher-restart/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'"

build: 

	go build -ldflags $(ldflag) -x .

docker:

	docker build -t {{.project}}:latest -f Dockerfile .

`
