# 太极

> 无极生太极，太极生两仪。。。
> 一个go脚手架,目的是能够编译出多平台的二进制,镜像,rpm,debu,等格式的软件包，自动生成版本且保持一致

## 安装使用

cd <项目目录>
GOPATH 下:
taiji init <文件类型>
非GOPATH 下:
taiji init <文件类型> --mod=<模块名字>

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

采用多阶编译,镜像中修改时区为国内，以及镜像源替换为国内的命令

### Makefile

在编译的时候注入版本信息到go文件中,如果有tag则为tag版本没有则为提交次数和hash,提供常用系统下的交编译命令,去除了字符链接缩小体积

## TODO

* 使用`git describe --tags --always --dirty="-dev"`替代if判断来生成自动生成版本号

* 多平台支持,本地直接生成已经是实现，docker，rpm,deb考虑通过buildkit实现
