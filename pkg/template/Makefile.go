package template

// Makefile 模板
const Makefile = `
# 判断当前commit是否有tag如果有tag则显示tag没有则显示 commit次数.哈希

v = $$(echo "$$(git log --oneline |wc -l).$$(git log -n1 --pretty=format:%h)" | sed 's/ //g')

tag = $(shell git log -n1 --pretty=format:%h |git tag --contains)

ifneq ($(tag),)

v = $$(git tag --sort=taggerdate |tail -1)

endif

path = $$(cat go.mod |grep module |cut -b 8-))

version=  -X '$(path)/version.Version=$(v)'

goversion = -X '$(path)/version.GoVersion=$$(go version | awk '{printf($$3)}')'

gitcommit = -X '$(path)/version.GitCommit=$$(git rev-parse HEAD)'

built = -X '$(path)/version.Built=$$(date "+%Y-%m-%d %H:%M:%S")'

ldflag = "-s -w $(version) $(goversion) $(gitcommit) $(built)"

# 编译为当前系统的二进制文件
.PHONY: build
build: 

	go build -ldflags $(ldflag)  -mod=vendor .
	
# 安装二进制文件到系统path
.PHONY: install
install:

	chmod +x {{.project}} && mv {{.project}} /usr/local/bin

# 编译为docker镜像
.PHONY: docker
docker:

	docker build -t {{.project}}:latest -f Dockerfile .

# 在容器中使用Makefile编译容器
build_in_docker:

	CGO_ENABLED=0 GOOS=linux go build -ldflags $(ldflag) -x -mod=vendor .

# 交叉编译为windows的二进制文件
.PHONY: windows
windows:

	GOOS=windows go build -ldflags $(ldflag) -x -mod=vendor .

# 交叉编译为苹果osx的二进制文件
.PHONY: darwin
darwin:

	GOOS=darwin go build -ldflags $(ldflag) -x -mod=vendor .

# 交叉编译为arm的linux环境（树莓派等环境）二进制文件
.PHONY: arm
arm:

	GOARCH=arm GOARM=7 GOOS=linux go build -ldflags $(ldflag) -x -mod=vendor .
`
