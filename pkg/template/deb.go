package template

func init() {
	Default.Add("deb", DEBDOCKERFILE, "build/deb/Dockerfile")
	Default.Add("debCtl", DEBCONTROL, "build/deb/common/control")
	Default.Add("debBuild", DEBDBUILD, "build/deb/build-deb")
}

const DEBDOCKERFILE = `#编译deb的
ARG GO_IMAGE
ARG BUILD_IMAGE

FROM ${GO_IMAGE} AS golang

FROM ${BUILD_IMAGE}

RUN echo "deb http://mirrors.aliyun.com/debian/ buster main non-free contrib"                       > /etc/apt/sources.list && \
    echo "deb-src http://mirrors.cloud.aliyuncs.com/debian/ buster main non-free contrib"           >>/etc/apt/sources.list && \
    echo "deb http://mirrors.cloud.aliyuncs.com/debian-security buster/updates main"                >>/etc/apt/sources.list && \
    echo "deb-src http://mirrors.cloud.aliyuncs.com/debian-security buster/updates main"            >>/etc/apt/sources.list && \
    echo "deb http://mirrors.cloud.aliyuncs.com/debian/ buster-updates main non-free contrib"       >>/etc/apt/sources.list && \
    echo "deb-src http://mirrors.cloud.aliyuncs.com/debian/ buster-updates main non-free contrib"   >>/etc/apt/sources.list && \
    echo "deb http://mirrors.cloud.aliyuncs.com/debian/ buster-backports main non-free contrib"     >>/etc/apt/sources.list && \
    echo "deb-src http://mirrors.cloud.aliyuncs.com/debian/ buster-backports main non-free contrib" >>/etc/apt/sources.list 

RUN apt-get update && apt-get install -y git make

ENV GOPROXY=direct
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

COPY --from=golang /usr/local/go /usr/local/go

ARG BUILD_DIR
COPY build/deb/build-deb/ /usr/local/bin/

RUN chmod +x /usr/local/bin/build-deb

WORKDIR /root/

ENTRYPOINT [ "build-deb" ]

`

const DEBCONTROL = `
Section: unknown
Priority: optional
Maintainer: naturelr naturelr@qq.com
Build-Depends: go
Homepage: https://github.com/NatureLR/taiji
Description: 打包测试
`
const DEBDBUILD = `#!/usr/bin/env bash

set -e 
set -x

ARCH=$(dpkg --print-architecture)

mkdir -p ~/debbuild/DEBIAN
mkdir -p ~/debbuild/usr/local/bin
cp /data/artifacts/tgz/$PROJECT-$VERSION.tar.gz .
cp -r /data/build/deb/common/* ~/debbuild/DEBIAN/

tar -xzf $PROJECT-$VERSION.tar.gz 
make -C $PROJECT build

cp $PROJECT/artifacts/bin/$PROJECT /root/debbuild/usr/local/bin

# systemd
if [ -d "/data/build/systemd" ];then 
mkdir -p ~/debbuild/usr/lib/systemd/system/
cp /data/build/systemd/* ~/debbuild/usr/lib/systemd/system/
fi

VER=$VERSION
if [[ $VERSION =~ v[0-9]+\.[0-9]+\.[0-9]+ ]]; then
  VER=$(echo "$VERSION" | sed 's/v//')
  echo "$VER"
fi

echo Source: $PROJECT >> ~/debbuild/DEBIAN/control
echo Package: $PROJECT >> ~/debbuild/DEBIAN/control
echo Version: $VER >> ~/debbuild/DEBIAN/control
echo Standards-Version: $VER >> ~/debbuild/DEBIAN/control
echo Architecture: $ARCH >> ~/debbuild/DEBIAN/control

dpkg-deb --build /root/debbuild/ $PROJECT-$VERSION-$ARCH.deb

mv $PROJECT-$VERSION-$ARCH.deb /data/artifacts/deb

`
