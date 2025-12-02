package template

func init() {
	Default.Add("deb", DEBDOCKERFILE, "build/deb/Dockerfile")
	Default.Add("debCtl", DEBCONTROL, "build/deb/control")
}

const DEBDOCKERFILE = `#编译deb的
ARG GO_IMAGE
ARG BUILD_IMAGE

FROM ${GO_IMAGE} AS golang

FROM ${BUILD_IMAGE} AS builder

ARG PROJECT
ARG VERSION

RUN apt-get update && apt-get install -y git make

ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

COPY --from=golang /usr/local/go /usr/local/go

WORKDIR /root/

RUN mkdir -p  /root/debbuild/DEBIAN  /root/debbuild/usr/local/bin  /root/debbuild/usr/lib/systemd/system/

COPY ./build/deb/control /root/debbuild/DEBIAN/control
COPY ./build/systemd/${PROJECT}.service ~/debbuild/usr/lib/systemd/system/
COPY ./artifacts/tgz/${PROJECT}-${VERSION}.tar.gz .

RUN tar -xzf ${PROJECT}-${VERSION}.tar.gz
RUN git config --global --add safe.directory /root/${PROJECT}
RUN make -C ${PROJECT} build

RUN cp ${PROJECT}/artifacts/bin/${PROJECT} /root/debbuild/usr/local/bin

RUN cat  /root/debbuild/DEBIAN/control

RUN sed -i "s/__source__/${PROJECT}/g"                         /root/debbuild/DEBIAN/control && \
    sed -i "s/__package__/${PROJECT}/g"                        /root/debbuild/DEBIAN/control && \
    sed -i "s/__version__/${VERSION#v}/g"                      /root/debbuild/DEBIAN/control && \
    sed -i "s/__standards_version__/${VERSION#v}/g"            /root/debbuild/DEBIAN/control && \
    sed -i "s/__architecture__/$(dpkg --print-architecture)/g" /root/debbuild/DEBIAN/control 

RUN cat /root/debbuild/DEBIAN/control
RUN dpkg-deb --build /root/debbuild/ ${PROJECT}-${VERSION}-$(dpkg --print-architecture).deb

FROM scratch AS export

COPY --from=builder /root/*deb .

`

const DEBCONTROL = `
Section: unknown
Priority: optional
Maintainer: youname <youemail.com>
Build-Depends: go
Homepage: <homepage>
Description: <Description>
Source: __source__
Package: __package__
Version:  __version__
Standards-Version: __standards_version__
Architecture:  __architecture__
`
