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
		fmt.Sprintf("build/rpm/%s.spec", filepath.Base(os.Getenv("PWD"))))
	Default.Add("rpmDockerfile", RPMDOCKERFILE, "build/rpm/Dockerfile")
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

{{.project}}/artifacts/bin/{{.project}} version

# 安装阶段需要做的
%install

install -D  -p  -m 0755 ${RPM_BUILD_DIR}/src/{{.project}}/artifacts/bin/{{.project}} ${RPM_BUILD_ROOT}%{_bindir}/{{.project}}
install -D -m 0644 ${RPM_BUILD_DIR}/src/{{.project}}/build/systemd/{{.project}}.service ${RPM_BUILD_ROOT}%{_unitdir}/{{.project}}.service

# 说明%{buildroot}中那些文件和目录需要打包到rpm中
%files

%{_bindir}/{{.project}}
%{_unitdir}/{{.project}}.service

# 变更记录
%changelog
`

// RPMDOCKERFILE 生成rpm包的dockerfile模板
const RPMDOCKERFILE = `# rpm构建环境
ARG GO_IMAGE
ARG BUILD_OS=centos
ARG BUILD_VERSION=8
ARG BUILD_IMAGE=${BUILD_OS}:${BUILD_VERSION}

# 从此镜像中获取go
FROM ${GO_IMAGE} AS golang

FROM ${BUILD_IMAGE}

# 更换源为阿里源
RUN mkdir /etc/yum.repos.d/bak && mv /etc/yum.repos.d/*repo /etc/yum.repos.d/bak && \
    curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-8.repo

# 安装rpmbuild 工具
RUN  yum install -y rpm-build rpmlint yum-utils rpmdevtools make git 
RUN  rpmdev-setuptree

WORKDIR /root/rpmbuild

# 从golang官方镜像中拷贝到centos镜像
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
COPY --from=golang /usr/local/go /usr/local/go

ENTRYPOINT ["/bin/rpmbuild"]

`
