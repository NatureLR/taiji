# 编译

FROM golang:1.13.5-alpine3.10 as build

WORKDIR /go/src/go-project/

COPY .  .

RUN pwd && ls

# 国内使用的goproxy
#RUN export GOPROXY=https://goproxy.cn

RUN  go build -o go-project -mod=vendor .

# 运行

FROM alpine:latest

WORKDIR /root/

COPY --from=build /go/src/go-project/go-project .


#EXPOSE port

#ENTRYPOINT ["./go-project"]