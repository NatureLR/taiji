package template

func init() {
	Default.Add("dockerfile", Dockerfile, "build/Dockerfile")
}

// Dockerfile 模版
const Dockerfile = `# 编译镜像 golang:x.y.z-alpine3.13
ARG BUILD_IMAGE=golang:1.18-alpine
ARG RUN_IMAGE=alpine:3

# 编译镜像
FROM ${BUILD_IMAGE} as build

ENV ROOT_DIR=/build
WORKDIR /build

COPY . .

# 修改源为国内阿里
# 修改时区为上海
# 安装make和git工具
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata  && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk add make && \
    apk add git

# 国内使用的goproxy
ENV GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct

RUN make build-in-docker

# 运行镜像
FROM ${RUN_IMAGE}

WORKDIR /root/

# 调整时区为北京时间
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#    apk add --no-cache ca-certificates tzdata  && \
#    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 添加nsswitch.conf，如不添加hosts文件无效
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

COPY --from=build /build/{{.project}} .

#EXPOSE <port>

#ENTRYPOINT ["./{{.project}}"]

CMD ["./{{.project}}"]
`
