# 编译

FROM golang:latest as build

WORKDIR /build

COPY .  .

# 国内使用的goproxy
#RUN export GOPROXY=https://goproxy.cn

RUN  CGO_ENABLED=0 GOOS=linux go build -o monitor -mod=vendor .
#RUN  CGO_ENABLED=0 GOOS=linux go build -o monitor .

# 运行

FROM alpine:latest

WORKDIR /root/

COPY --from=build /build/monitor .


#EXPOSE port

ENTRYPOINT ["./go-project"]