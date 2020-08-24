package template

func init() {
	DefaultPool.add("dockerfile", Dockerfile, "Dockerfile")
}

// Dockerfile 模版
const Dockerfile = `
# 编译镜像
FROM golang:1.15-alpine as build

WORKDIR /build

COPY .  .

# 修改源为国内阿里
# 修改时区为上海
# 安装make和git工具
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata  && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk add make && \
    apk add git

# 国内使用的goproxy
#ENV GOPROXY=https://goproxy.cn

RUN make build_in_docker

# 运行镜像
FROM alpine:latest

WORKDIR /root/

# 调整时区为北京时间
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#    apk add --no-cache ca-certificates tzdata  && \
#    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=build /build/{{.project}} .

#EXPOSE <port>

#ENTRYPOINT ["./{{.project}}"]

CMD ["./{{.project}}"]
`
