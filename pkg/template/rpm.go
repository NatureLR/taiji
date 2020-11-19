package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	Default.Add(
		"rpm",
		RPMSPEC,
		fmt.Sprintf("build/rpm/SPECS/%s.spec", filepath.Base(os.Getenv("PWD"))))
	Default.Add("rpmMakefile", RPMMAKEFILE, "build/rpm/Makefile")
}

// RPMSPEC 生成rpm包的spec模板
const RPMSPEC = `# 构建rpm包的文件
%global debug_package %{nil}

Name:           {{.project}}
Version:        %{_version}
Release:        1%{?dist}
Summary:        {{.ShortDescribe}}

Group:          Application/WebServer
License:        Apache 2.0
URL:            http://www.baidu.com
Source0:        %{name}.tar.gz

# 构建依赖
BuildRequires:  git
BuildRequires:  make

# 详细描述
%description

{{.LongDescribe}}

# 构建之前执行的脚本，一般为解压缩
%prep

# %setup 不加任何选项，仅将软件包打开。
# %setup -a 切换目录前，解压指定 Source 文件，例如 "-a 0" 表示解压 "Source0"
# %setup -n newdir 将软件包解压在newdir目录。
# %setup -c 解压缩之前先产生目录。
# %setup -b num 将第 num 个 source 文件解压缩。
# %setup -D 解压前不删除目录
# %setup -T 不使用default的解压缩操作。
# %setup -q 不显示解包过程
# %setup -T -b 0 将第 0 个源代码文件解压缩。
# %setup -c -n newdir 指定目录名称 newdir，并在此目录产生 rpm 套件。
%setup -q -c -n src -a 0

# 编译脚本
%build

cd {{.project}} && make build

# 检查
%check

{{.project}}/bin/{{.project}} version

# 安装阶段需要做的
%install

install -D  -p  -m 0755 ${RPM_BUILD_DIR}/src/{{.project}}/bin/{{.project}} ${RPM_BUILD_ROOT}%{_bindir}/{{.project}}
install -D -m 0644 ${RPM_BUILD_DIR}/src/{{.project}}/build/systemd/{{.project}}.service ${RPM_BUILD_ROOT}%{_unitdir}/{{.project}}.service

# 说明%{buildroot}中那些文件和目录需要打包到rpm中
%files

%{_bindir}/{{.project}}
%{_unitdir}/{{.project}}.service

# 变更记录
%changelog
`

// RPMMAKEFILE 生成rpm包的makefile模板
const RPMMAKEFILE = `# 通过容器构建各个系统各个版本的rpm包
VERSION ?= $(shell git describe --tags --always --dirty="-dev" )
# ================ go版本配置 ================
GO_VERSION ?= 1.15
GO_BASE_IMAGE ?= golang
GO_IMAGE ?= $(GO_BASE_IMAGE):$(GO_VERSION)

# 将项目打包的tgz文件放入rpmbuild/SOURCES
tgz ?= mkdir -p rpmbuild/SOURCES  && if [ ! -d "../tgz" ]; then echo tgz文件不存在创建tgz包;$(MAKE) -C ../ tgz && cp -f ../tgz/*tar.gz rpmbuild/SOURCES;fi

# 根据各个系统构建编译环境的容器
BUILD ?= DOCKER_BUILDKIT=1 \
	docker build \
	$(BUILD_IMAGE_FLAG) \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	-t rpmbuild-$@ \
	-f $@/Dockerfile \
	.

# ================ 配置构建的specs文件 ================
SPEC_FILES ?= {{.project}}.spec 
SPECS ?= $(addprefix SPECS/, $(SPEC_FILES))
# 在各个系统
RPMBUILD_FLAG ?= -ba \
	--define '_version ${VERSION}' \
	$(SPECS)
# 在容器里运行rpmbuild打包生成rpm文件
RUN ?= docker run \
	-v $(CURDIR)/rpmbuild/RPMS:/root/rpmbuild/RPMS \
	-v $(CURDIR)/rpmbuild/SRPMS:/root/rpmbuild/SRPMS \
	-v $(CURDIR)/rpmbuild/SOURCES:/root/rpmbuild/SOURCES \
	rpmbuild-$@ $(RPMBUILD_FLAG)

# ================ 目标操作系统配置 ================
CENTOS_RELEASES ?= centos-8
DISTROS := $(CENTOS_RELEASES)

.PHONY: help
help: ## 显示make的目标
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## 删除rpmbuild的包和中间产生的文件
	rm -rf rpmbuild

.PHONY: $(DISTROS)
$(DISTROS):
	@echo "================ 构建$@ ================"
	$(tgz)
	$(BUILD)
	$(RUN)
	@echo "================ 构建$@完成 ================"

.PHONY: rpm
rpm: centos ## 构建使用rpm系统的包
	
.PHONY: centos
centos: $(CENTOS_RELEASES) ## 构建centos的rpm包

.PHONY: centos-8
centos-8: ## 构建centos8的rpm包
`
