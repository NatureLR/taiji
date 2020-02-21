# 编译

FROM golang:1.13.6 as build

WORKDIR /go/src/go-project/

COPY .  .

# 国内使用的goproxy
#RUN export GOPROXY=https://goproxy.cn

RUN  make build_in_docker

# 运行

FROM alpine:latest

WORKDIR /root/

COPY --from=build /go/src/go-project/go-project .


#EXPOSE port

#ENTRYPOINT ["./go-project"]