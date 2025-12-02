package template

func init() {
	Default.Add("makefile", Makefile, "Makefile")
	Default.Add("buildmakefile", BuildMakefile, "build/common.mk")
}

// Makefile 模板
const Makefile = `# 主文件
include ./build/common.mk

##@ General

.PHONY: help
help: ## 显示make帮助
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## 清理编译输出
	@rm -rf $(OUTPUT_DIR)

.PHONY: swag
swag: ## 生成swagger文档
	@swag init --parseDependency --parseInternal
	
.PHONY: doc
doc: swag ## 生成swagger文档和
	@go run . doc

##@ Build

.PHONY: build
build: ## 本地编译当前系统和架构
	@echo $(GREEN)编译$(GOOS)/$(GOARCH)
	@$(BUILD)
	@cp $(GO_OUTPUT) $(OUTPUT_BIN_DIR)/$(PROJECT)

.PHONY: run
run: ## 直接运行不编译
	go run $(ROOT_DIR)

.PHONY: build-all
build-all: ## 多平台多架构
	@for os in $(OSS);do \
		for arch in $(ARCHS);do \
			GOOS=$$os GOARCH=$$arch make build; \
		done \
	done

.PHONY: build-in-docker
build-in-docker: ## 在docker里的编译选项
	@CGO_ENABLED=0 go build -ldflags $(LDFLAG) -o $(PROJECT) $(ROOT_DIR)

all: build-all docker tgz rpm deb ## 编译打包所有

##@ Deploy
.PHONY: install
install: build ## 安装到系统PATH
	@cp $(GO_OUTPUT) /usr/local/bin/$(PROJECT)

.PHONY: uninstall
uninstall: ## 卸载安装
	@rm -rf /usr/local/bin/$(PROJECT)

##@ Packag

.PHONY: docker
docker: ## 编译docker镜像
	@echo $(GREEN)构建docker镜像
	@$(DOCKER_BUILD)
	@echo $(YELLOW)镜像地址:
	@echo $(IMAGE_ADDR)
	@echo $(IMAGE_ADDR_LATEST)

.PHONY: tgz
tgz: ## 打包为tar包
	@echo $(GREEN)打包为tgz
	@$(TGZ)

.PHONY: rpm
rpm: ## 打包为rpm包
	@echo $(GREEN)打包rpm
	@mkdir -p $(OUTPUT_RPM_DIR)/RPMS $(OUTPUT_RPM_DIR)/SRPMS
	@$(CHECK_TGZ)
	@$(RPM_DOCKER_BUILD)

.PHONY: deb
deb: ## 打包为deb包
	@echo $(GREEN)打包deb
	@mkdir -p $(OUTPUT_DEB_DIR)
	@$(CHECK_TGZ)
	@$(DEB_DOCKER_BUILD)
	@$(DEB_DOCKER_RUN)

`

// BuildMakefile build目录下的makefile需要和根目录下的makefile合并
const BuildMakefile = `# 全局配置
PROJECT := {{.project}}
ARCHS   := amd64 arm64
OSS     := linux windows darwin

# 判断当前commit是否有tag如果有tag则显示tag没有则显示自动生成的tag
# 版本号格式说明:
# 如v0.1.1-dev.1.5.097ce20，表示最近的tag是v0.1.1，距离最近的tag是1个提交，一共有5个commit，commit hash是097ce20
ZERO_TAG     := v0.0.0
DIRTY        ?=dev
# 距离最近的tag的提交数量
SINCE_TAG    := $(shell echo $(shell git rev-list $(shell git describe --tags --abbrev=0)..HEAD --count))
# 总的commit数量
COUNT_COMMIT := $(shell git log --oneline |wc -l)
# 最近的commit hash
COMMIT_HASH  := $(shell git log -n1 --pretty=format:%h)
COUNT_TAG	 := $(shell git tag --list |wc -l)
# 最近的tag
LAST_TAG     =  $(shell git describe --tags --always --dirty=-$(DIRTY))
# 如果一个tag都没有则使用默认的tag
ifeq ($(strip $(COUNT_TAG)),0)
LAST_TAG     =  $(shell echo "$(ZERO_TAG)-$(DIRTY)")
endif
AUTO_VERSION =  $(shell echo "$(LAST_TAG).$(SINCE_TAG).$(COUNT_COMMIT).$(COMMIT_HASH)" | sed 's/ //g')
VERSION      =  $(AUTO_VERSION)
TAG          =  $(shell git log -n1 --pretty=format:%h |git tag --contains)
ifneq ($(TAG),)
VERSION      =  $(shell git tag --sort=committerdate |tail -1)
endif

# go 参数
GOOS       ?= $(shell go env GOOS)
GOARCH     ?= $(shell go env GOARCH)
# 使用本地go版本作为go版本
# GOVERSION  ?= $(shell go version | grep -Eo '[0-9]+\.[0-9]+\.[0-9]+')
GOVERSION  ?= {{.GoVersion}}

# 目录
ROOT_DIR          := $(realpath $(CURDIR))
BUILD_DIR         := $(ROOT_DIR)/build
BUILD_RPM_DIR     := $(BUILD_DIR)/rpm
INSTALL_DIR       := /usr/local/bin
OUTPUT            := artifacts
OUTPUT_DIR        := $(ROOT_DIR)/$(OUTPUT)
OUTPUT_BIN_DIR    := $(OUTPUT_DIR)/bin
OUTPUT_RPM_DIR    := $(OUTPUT_DIR)/rpm
OUTPUT_DEB_DIR    := $(OUTPUT_DIR)/deb
OUTPUT_TGZ_DIR    := $(OUTPUT_DIR)/tgz
PRJVER            := $(PROJECT)-$(VERSION)

# go 注入参数
GO_PATH      := $(shell cat $(ROOT_DIR)/go.mod |grep module |cut -b 8-)
X_VERSION    := -X '$(GO_PATH)/pkg/versions.xVersion=$(VERSION)'
X_GIT_COMMIT := -X '$(GO_PATH)/pkg/versions.xGitCommit=$$(git rev-parse HEAD)'
X_BUILT      := -X '$(GO_PATH)/pkg/versions.xBuilt=$$(date "+%Y-%m-%d %H:%M:%S")'
LDFLAG       := "-s -w $(X_VERSION) $(X_GIT_COMMIT) $(X_BUILT)"
GO_OUTPUT    := $(OUTPUT_BIN_DIR)/$(PRJVER)-$(GOOS)-$(GOARCH)
ifeq ($(GOOS),windows)
GO_OUTPUT    := $(OUTPUT_BIN_DIR)/$(PRJVER)-$(GOOS)-$(GOARCH).exe
endif
BUILD        := go build -ldflags $(LDFLAG) -o $(GO_OUTPUT) $(ROOT_DIR)

TGZ_CMD   := tar --exclude $(PROJECT)/$(OUTPUT) -czf $(PRJVER).tar.gz $(PROJECT)
TGZ       := mkdir -p $(OUTPUT_TGZ_DIR) && cd $(CURDIR)/../ && $(TGZ_CMD) &&  mv $(PRJVER).tar.gz $(OUTPUT_TGZ_DIR)
CHECK_TGZ := if [ ! -f "$(OUTPUT_TGZ_DIR)/$(PRJVER).tar.gz" ]; then echo tgz文件不存在创建tgz包;$(MAKE) tgz;fi

# docker
GO_IMAGE         ?= golang:$(GOVERSION)-trixie
# 产生镜像时用于运行的镜像
GO_RUN_IMAGE     ?= alpine:latest
GO_BUILD_IMAGE   ?= golang:$(GOVERSION)-alpine
GO_BASE_IMAGE    ?= golang:$(GOVERSION)
RPM_BUILD_IMAGE  ?= rockylinux:9
DEB_BUILD_IMAGE  ?= debian:trixie
PLATFORM         ?= linux/amd64,linux/arm64
DOCKER_PLATFORM  ?= $(PLATFORM)
RPM_PLATFORM     ?= $(PLATFORM)
DEB_PLATFORM     ?= $(PLATFORM)

# 自己的仓库
DOCKER_REPO       = naturelr
IMAGE_ADDR        = $(DOCKER_REPO)/$(PROJECT):$(VERSION)
IMAGE_ADDR_LATEST = $(DOCKER_REPO)/$(PROJECT):latest
ifeq ($(DOCKER_REPO),)
IMAGE_ADDR        = $(PROJECT):$(VERSION)
IMAGE_ADDR_LATEST = $(PROJECT):latest
endif

DOCKER_BUILD     := docker buildx build \
	--platform $(DOCKER_PLATFORM) \
	-t $(IMAGE_ADDR) \
	-t $(IMAGE_ADDR_LATEST) \
	--build-arg RUN_IMAGE=$(GO_RUN_IMAGE) \
	--build-arg BUILD_IMAGE=$(GO_BUILD_IMAGE) \
	-f $(BUILD_DIR)/Dockerfile \
	-o type=registry \
	$(ROOT_DIR)
RPM_DOCKER_BUILD := docker buildx build \
	--platform $(RPM_PLATFORM) \
	-f $(BUILD_DIR)/rpm/Dockerfile \
	--output type=local,dest=$(OUTPUT_RPM_DIR) \
	--build-arg VERSION=$(VERSION) \
	--build-arg GO_IMAGE=$(GO_BASE_IMAGE) \
	--build-arg BUILD_IMAGE=$(RPM_BUILD_IMAGE) \
	--build-arg PROJECT=$(PROJECT) \
	$(ROOT_DIR)
DEB_DOCKER_BUILD := docker buildx build \
	--platform $(DEB_PLATFORM) \
	-f $(BUILD_DIR)/deb/Dockerfile \
	--output type=local,dest=$(OUTPUT_DEB_DIR) \
	--build-arg GO_IMAGE=$(GO_BASE_IMAGE) \
	--build-arg BUILD_DIR=$(DEB_DIR) \
	--build-arg BUILD_IMAGE=$(DEB_BUILD_IMAGE)\
	--build-arg PROJECT=$(PROJECT) \
	--build-arg VERSION=$(VERSION) \
	--progress=plain \
	--no-cache \
	$(ROOT_DIR)

# 颜色
RED    := $(shell tput -Txterm setaf 1)
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
VIOLET := $(shell tput -Txterm setaf 5)
AQUA   := $(shell tput -Txterm setaf 6)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
`
