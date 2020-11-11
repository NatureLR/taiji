package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	Default.Add(
		"systemd",
		Systemd,
		fmt.Sprintf("build/systemd/%s.service", filepath.Base(os.Getenv("PWD"))))
}

// Systemd 模板
const Systemd = `
# 启动顺序与依赖关系
[Unit]
# 简单的描述
Description={{.ShortDescribe}}
# 文档地址
#Documentation=https://github.com/NatureLingRan/go-project
# 在某服务启动之前启动
#Before=
# multi-user.target组之后启动
After=multi-user.target
# 弱依赖network-online.target这个组
Wants=network-online.target

# 启动行为
[Service]
# 启动类型
Type=simple
# 环境变量
#Environment=
# 环境变量文件
#EnvironmentFile=

# 定义启动执行的命令
ExecStart=/usr/local/bin/{{.project}}
# 定义重启执行的命令
ExecReload=/bin/kill -s HUP $MAINPID
# 重启策略
Restart=always
# 重启服务需要等待的秒数
RestartSec=3

#定义如何安装这个配置文件，即怎样做到开机启动。
[Install]
WantedBy=multi-user.target
`
