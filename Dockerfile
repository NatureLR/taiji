# 编译镜像
FROM golang:1.14-alpine as build

WORKDIR /build

COPY .  .

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

COPY --from=build /build/go-project .

#EXPOSE <port>

#ENTRYPOINT ["./go-project"]

CMD ["./go-project"]
