package template

func init() {
	Default.Add("makefile", Makefile, "Makefile")
	Default.Add("BUILDMAKEFILE", BUILDMAKEFILE, "build/Makefile")
}

// Makefile 模板
const Makefile = `# 判断当前commit是否有tag如果有tag则显示tag没有则显示 commit次数.哈希
VER = $(shell echo "$(shell git log --oneline |wc -l).$(shell git log -n1 --pretty=format:%h)" | sed 's/ //g')
# 用于判断当前的提前是否有tag
TAG = $(shell git log -n1 --pretty=format:%h |git tag --contains)
# 如果没有手动指定标签就使用自动生成的标签
ifneq ($(TAG),)
# 通过git tag命令指定的标签
VER = $(shell git tag --sort=committerdate |tail -1)
endif

VERSION =  -X '$(GO_PATH)/pkg/versions.xVersion=$(VER)'
GO_PATH = $(shell cat go.mod |grep module |cut -b 8-)
GIT_COMMIT = -X '$(GO_PATH)/pkg/versions.xGitCommit=$$(git rev-parse HEAD)'
BUILT = -X '$(GO_PATH)/pkg/versions.xBuilt=$$(date "+%Y-%m-%d %H:%M:%S")'
LDFLAG = "-s -w $(VERSION) $(GO_VERSION) $(GIT_COMMIT) $(BUILT)"

PROJECT = {{.project}}
# 二进制文件生成目录
BIN_DIR = bin
BUILD = go build -ldflags $(LDFLAG) -o $(BIN_DIR)/$(PROJECT) .

.PHONY: help
help: ## 显示make的目标
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## 清理
	rm -rf bin
	$(MAKE) -C build clean

.PHONY: build
build: ## 编译为当前系统的二进制文件
	@$(BUILD)

.PHONY: run
run: ## 直接运行不编译
	@go run .

.PHONY: install
install: ## 安装二进制文件到系统path
	@chmod +x $(BIN_DIR)/$(PROJECT) && mv $(BIN_DIR)/$(PROJECT) /usr/local/bin

.PHONY: docker
ifeq ($(REGISTRY),)
docker: ## 编译为docker镜像
	@docker build -t $(PROJECT):latest -t $(PROJECT):$(VER) -f Dockerfile .
else
docker: ## 编译为docker镜像 
	@docker build -t $(REGISTRY)/$(PROJECT):latest -t $(REGISTRY)/$(PROJECT):$(VER) -f Dockerfile .
endif

.PHONY: build-in-docker
build-in-docker: ## Dockerfile中执行编译
	@CGO_ENABLED=0 GOOS=linux go build -ldflags $(LDFLAG) .

.PHONY: windows
windows: ## 交叉编译为windows的二进制文件
	@GOOS=windows $(BUILD)

.PHONY: darwin
darwin: ## 交叉编译为苹果osx的二进制文件
	@GOOS=darwin $(BUILD)

.PHONY: arm
arm: ## 交叉编译为arm的linux环境（树莓派等环境）二进制文件
	@GOARCH=arm GOARM=7 GOOS=linux $(BUILD)
`

// BUILDMAKEFILE build目录下的makefile需要和根目录下的makefile合并
const BUILDMAKEFILE = `# 编译rpm包和tar包需要和根目录下的Makefil合并
PROJECT := {{.project}}
ROOT_DIR ?= $(realpath $(CURDIR)/../)

# 判断当前commit是否有tag如果有tag则显示tag没有则显示 commit次数.哈希
VER = $(shell echo "$(shell git log --oneline |wc -l).$(shell git log -n1 --pretty=format:%h)" | sed 's/ //g')
# 用于判断当前的提前是否有tag
TAG = $(shell git log -n1 --pretty=format:%h |git tag --contains)
# 如果没有手动指定标签就使用自动生成的标签
ifneq ($(TAG),)
# 通过git tag命令指定的标签
VER = $(shell git tag --sort=committerdate |tail -1)
endif

REGISTRY ?=
BUILD_IMAGE:=docker build \
			-t $(REGISTRY)$(PROJECT):latest \
			-t $(REGISTRY)$(PROJECT):$(VER) \
			-f $(ROOT_DIR)/build/Dockerfile \
			$(ROOT_DIR)

GO_PATH ?= $(shell cat $(ROOT_DIR)/go.mod |grep module |cut -b 8-)
VERSION ?= -X '$(GO_PATH)/pkg/versions.xVersion=$(VER)'
GIT_COMMIT ?= -X '$(GO_PATH)/pkg/versions.xGitCommit=$$(git rev-parse HEAD)'
BUILT ?= -X '$(GO_PATH)/pkg/versions.xBuilt=$$(date "+%Y-%m-%d %H:%M:%S")'
LDFLAG ?= -s -w \
		$(VERSION) \
		$(GO_VERSION) \
		$(GIT_COMMIT) \
		$(BUILT)
# 二进制文件生成目录
BIN_DIR = $(ROOT_DIR)/bin
BUILD = go build -ldflags "$(LDFLAG)" -o $(BIN_DIR)/$(PROJECT) $(ROOT_DIR)

.PHONY: help
help: ## 显示make的目标
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## 清理构建生产的文件
	rm -rf $(ROOT_DIR)/bin
	$(MAKE) -C rpm clean
	rm -rf tgz

.PHONY: all
all: build rpm docker  ## 编译和构建所有支持的选项

.PHONY: tgz 
tgz: ## 将项目打包为tar.gz
	mkdir -p $(ROOT_DIR)/build/tgz
	docker run --rm \
		-w / \
		-v $(realpath $(ROOT_DIR)/../$(PROJECT)):/$(PROJECT) \
		-v $(ROOT_DIR)/build/tgz:/v \
		alpine \
		sh -c "tar -cf $(PROJECT).tar.gz $(PROJECT) && mv $(PROJECT).tar.gz /v"		

.PHONY: run
run: ## 直接运行不编译
	go run $(ROOT_DIR)

#  ================ 直接使用系统的go进行编译 ================
.PHONY: build
build: ## 编译为当前系统的二进制文件
	@echo ================ 本地构建可执行文件 ================
	$(BUILD)

.PHONY: install
install: ## 安装二进制文件到系统path
	chmod +x $(BIN_DIR)/$(PROJECT) && mv $(BIN_DIR)/$(PROJECT) /usr/local/bin

# TODO 各个架构的镜像
#  ================ 构建docker镜像 ================
.PHONY: docker
docker: ## 编译为docker镜像
	$(BUILD_IMAGE)

.PHONY: build-in-docker
build-in-docker: ## Dockerfile中执行编译
	CGO_ENABLED=0 GOOS=linux $(BUILD)

# TODO centos 7等rpm支持的包
#  ================ 使用容器构建rpm包 ================
.PHONY: rpm	
rpm: ## 构建rpm系统的包
	@echo ================ 容器构建RPM包 ================
	$(MAKE) -C ./rpm rpm 

.PHONY: centos	
centos: ## 构建centos系统
	$(MAKE) -C ./rpm centos

.PHONY: centos-8	
entos-8: ## 构建centos-8系统的包
	$(MAKE) -C ./rpm centos-8

`
