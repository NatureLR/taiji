# GO-PROJECT

    一个创建GO项目所需要的文件目录的程序,目前可以创建makefile,dockerfile,.gitignore

文件或目录|说明|实现程度
-|-|-
makefile  |用于编译的脚本  |实现了go各大部分平台的交叉编译,docker,以及去符号版本注入      |
dockerfile|编译为容器的文件 |编译镜像使用golang的完整版镜像,运行使用`alpine`镜像以减小体积|
.gitignore|让git忽略的文件 |忽略项目本身编译的二进制文件                              |
k8s.yaml  |k8s资源文件    |程序的Deployment,和service,镜像地址需要自己填写            |
version   |版本相关的代码  |通过命令在编译的时候注入版本信息,已实现编译时间;commit版本号等 |
cmd       |命令行参数代码  |croba的命令行代码                                       |
pkg       |库代码         |-                                                     |

## 安装使用

    go install github.com/NatureLingRan/go-project
    cd /<项目目录>
    go-project init

## 注意

    非go mod下版本无法注入
