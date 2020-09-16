# GO-PROJECT

一个创建GO项目所需要的文件目录的程序如下


## 安装使用

cd <项目目录>
GOPATH 下:
go-project init <文件类型>
非GOPATH 下:
go-project init <文件类型> --mod=<模块名字>

## 文件说明

### 命令行

* 命令使用[cobra](https://github.com/spf13/cobra)详情可以查看官方文档,创建的cmd文件夹即为`cobra`的命令入口

### 配置文件

配置使用[viper](https://github.com/spf13/viper),读取配置文件名字为`config.yaml`;
默认会读取以下目录

* 程序的根目录

* 程序下的config目录

* /etc/<程序的名字>目录

* home目录中的config.yaml

### Dockerfile

采用多阶编译,默认使用编译镜像使用`golang:1.15-alpine`,运行镜像为`alpine:latest`,可自行替换私有仓库
镜像中修改时区为国内，以及镜像源替换为国内的命令

### Makefile

在编译的时候注入版本信息到go文件中,如果有tag则为tag版本没有则为提交次数和hash,提供常用系统下的交编译命令,去除了字符链接缩小体积

### K8s

描述k8s资源的文件，需要自行替换镜像

### gitignore

git屏蔽的文件，默认有项目的名字和.vscode文件夹

### License

默认为apache2.0协议
