package template

// Dockerfile 模版
const Dockerfile = `
# 编译

FROM golang:1.13.5-alpine3.10 as build

WORKDIR /build

COPY .  .

# 国内使用的goproxy
#RUN export GOPROXY=https://goproxy.cn

RUN go build -o {{.project}} -mod=vendor .
#RUN  CGO_ENABLED=0 GOOS=linux go build -o {{.project}} .

# 运行

FROM alpine:latest

WORKDIR /root/

COPY --from=build /build/{{.project}} .


#EXPOSE <port>

ENTRYPOINT ["./{{.project}}"]

`
