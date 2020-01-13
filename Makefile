project = $$(go list)

version = -X '$(project)/version.Version=$$(git tag --sort=taggerdate |tail -1)'

goversion = -X '$(project)/version.GoVersion=$$(go version | awk '{printf($$3)}')'

gitcommit = -X '$(project)/version.GitCommit=$$(git rev-parse HEAD)'

built = -X '$(project)/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'

ldflag = "-s -w $(version) $(goversion) $(gitcommit) $(built)"

build: 

	go build -ldflags $(ldflag) -x .

docker:

	docker build -t go-project:latest -f Dockerfile .
