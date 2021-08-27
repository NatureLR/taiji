# 项目的名字
PROJECT = test

# go 参数
GO_OS      =
GO_ARCH    =
GO_VERSION =

# go 注入参数
GO_PATH    = $(shell cat go.mod |grep module |cut -b 8-)
VERSION    = -X '$(GO_PATH)/pkg/versions.xVersion=$(VER)'
GIT_COMMIT = -X '$(GO_PATH)/pkg/versions.xGitCommit=$$(git rev-parse HEAD)'
BUILT      = -X '$(GO_PATH)/pkg/versions.xBuilt=$$(date "+%Y-%m-%d %H:%M:%S")'
LDFLAG     = "-s -w $(VERSION) $(GO_VERSION) $(GIT_COMMIT) $(BUILT)"
BUILD      = go build -ldflags $(LDFLAG) -o $(BIN_DIR)/$(PROJECT) .

# 颜色
RED    := $(shell tput -Txterm setaf 1)
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
VIOLET := $(shell tput -Txterm setaf 5)
AQUA   := $(shell tput -Txterm setaf 6)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

# docker
DOCKER_REPO = hub.docker.com

# 目录
ROOT_DIR = 
BIN_DIR  = artifacts/$(GO_ARCH)/bin
RPM_DIR  = artifacts/$(GO_ARCH)/rpm
DEB_DIR  = artifacts/$(GO_ARCH)/deb
