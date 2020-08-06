# GO-PROJECT

    一个创建GO项目所需要的文件目录的程序,目前可以创建makefile,dockerfile,.gitignore

文件或目录|说明|说明
-|-|-
makefile  |用于编译的脚本  |编译文件，版本注入 |
dockerfile|编译为容器的文件 |编译doker镜像    |
.gitignore|让git忽略的文件 |git忽略文件       |
k8s.yaml  |k8s资源文件    |k8s部署文件       |
version   |版本相关的代码  |版本相关          |
cmd       |命令行参数代码  |croba的命令行代码  |
pkg       |库代码         |-                |

## 安装使用

    go install github.com/NatureLingRan/go-project
    cd <项目目录>
    go-project init

## 创建出的文件

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

采用多阶编译,默认使用编译镜像使用`golang:1.14-alpine`,运行镜像为`alpine:latest`,可自行替换私有仓库
镜像中修改时区为国内，以及镜像源替换为国内的命令

### Makefile

在编译的时候注入版本信息到go文件中,如果有tag则为tag版本没有则为提交次数和hash,提供常用系统下的交编译命令,去除了字符链接缩小体积

### K8s

描述k8s资源的文件，需要自行替换镜像

### gitignore

git屏蔽的文件，默认有项目的名字和.vscode文件夹

### License

默认为apache2.0协议
