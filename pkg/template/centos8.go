package template

func init() {
	Default.Add("centos8", Centos8, "build/rpm/centos-8/Dockerfile")
}

// Centos8 rpm打包dockerfile
const Centos8 = `
ARG GO_IMAGE
ARG BUILD_OS=centos
ARG BUILD_VERSION=8
ARG BUILD_IMAGE=${BUILD_OS}:${BUILD_VERSION}

FROM ${GO_IMAGE} AS golang

FROM ${BUILD_IMAGE}

# 更换源为阿里源
RUN mkdir /etc/yum.repos.d/bak && mv /etc/yum.repos.d/*repo /etc/yum.repos.d/bak && \
    curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-8.repo

# 安装rpmbuild 工具
RUN  yum install -y rpm-build rpmlint yum-utils rpmdevtools make git 
RUN  rpmdev-setuptree

WORKDIR /root/rpmbuild

COPY SPECS SPECS

# 从golang官方镜像中拷贝到centos镜像
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
COPY --from=golang /usr/local/go /usr/local/go

ENTRYPOINT ["/bin/rpmbuild"]
`
