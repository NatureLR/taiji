# 判断当前commit是否有tag如果有tag则显示tag没有则显示 commit次数.哈希
VER = $(shell echo "$(shell git log --oneline |wc -l).$(shell git log -n1 --pretty=format:%h)" | sed 's/ //g')

# 用于判断当前的提前是否有tag
TAG = $(shell git log -n1 --pretty=format:%h |git tag --contains)

# 如果没有手动指定标签就使用自动生成的标签
ifneq ($(TAG),)

# 手动指定的标签
VER = $(shell git tag --sort=committerdate |tail -1)

endif

GO_PATH = $(shell cat go.mod |grep module |cut -b 8-)

VERSION =  -X '$(GO_PATH)/cmd.version=$(VER)'

GO_VERSION = -X '$(GO_PATH)/cmd.goVersion=$$(go version | awk '{printf($$3)}')'

GIT_COMMIT = -X '$(GO_PATH)/cmd.gitCommit=$$(git rev-parse HEAD)'

BUILT = -X '$(GO_PATH)/cmd.built=$$(date "+%Y-%m-%d %H:%M:%S")'

LDFLAG = "-s -w $(VERSION) $(GO_VERSION) $(GIT_COMMIT) $(BUILT)"

PROJECT = go-project

# 二进制文件生成目录
BIN_DIR = bin

BUILD = go build -ldflags $(LDFLAG) -o $(BIN_DIR)/$(PROJECT) .

# 编译为当前系统的二进制文件
.PHONY: build
build: 

	@$(BUILD)

# 直接运行不编译
.PHONY: run
run:
	
	@go run .

# 安装二进制文件到系统path
.PHONY: install
install:

	@chmod +x $(BIN_DIR)/$(PROJECT) && mv $(BIN_DIR)/$(PROJECT) /usr/local/bin

# 清理
.PHONY: clean
clean:
	
	@rm -rf bin	test

# 编译为docker镜像
.PHONY: docker
ifeq ($(REGISTRY),)
docker:
	@docker build -t $(PROJECT):latest -t $(PROJECT):$(VER) -f Dockerfile .
else
docker: 
	@docker build -t $(REGISTRY)/$(PROJECT):latest -t $(REGISTRY)/$(PROJECT):$(VER) -f Dockerfile .
endif

# Dockerfile中执行编译
.PHONY: build_in_docker
build_in_docker:

	@CGO_ENABLED=0 GOOS=linux go build -ldflags $(LDFLAG) .

# 交叉编译为windows的二进制文件
.PHONY: windows
windows:

	@GOOS=windows $(BUILD)

# 交叉编译为苹果osx的二进制文件
.PHONY: darwin
darwin:

	@GOOS=darwin $(BUILD)

# 交叉编译为arm的linux环境（树莓派等环境）二进制文件
.PHONY: arm
arm:

	@GOARCH=arm GOARM=7 GOOS=linux $(BUILD)
