# 编译

FROM golang:latest as build

WORKDIR /build

COPY .  .

# 国内使用的goproxy
#RUN export GOPROXY=https://goproxy.cn

RUN  CGO_ENABLED=0 GOOS=linux go build -o go-project -mod=vendor .
#RUN  CGO_ENABLED=0 GOOS=linux go build -o go-project .

# 运行

FROM alpine:latest

WORKDIR /root/

COPY --from=build /build/go-project .


#EXPOSE port

ENTRYPOINT ["./go-project"]