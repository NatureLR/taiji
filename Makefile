path = $$(go list)

version = -X '$(path)/version.Version=$$(git tag --sort=taggerdate |tail -1)'

goversion = -X '$(path)/version.GoVersion=$$(go version | awk '{printf($$3)}')'

gitcommit = -X '$(path)/version.GitCommit=$$(git rev-parse HEAD)'

built = -X '$(path)/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'

ldflag = "-s -w $(version) $(goversion) $(gitcommit) $(built)"

build: 

	go build -ldflags $(ldflag) -x .

docker:

	docker build -t go-project:latest -f Dockerfile .
