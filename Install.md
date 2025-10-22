# 专治各种安装环境

## Go
安装 `Go` 有几个点注意下：
* `Go` 语言本身是自编译语言，在一个没有 `Go` 环境的机器上直接安装 `gvm` 会出问题
* 虽然 `gvm` 会解决各个版本的切换问题，但是最佳实践还是先安装一个版本的 `Go`
* 直接去官方找安装资源

### 1. 使用版本工具安装

`Ubuntu` 环境
```shell
# 先安装一个 Go 环境
# 笨方法
wget https://golang.google.cn/dl/go1.25.3.linux-arm64.tar.gz

# 如果不知道自己本地环境是什么类型的 是 arm 还是 amd 使用如下方式
dog@dog:~/install$ uname -m
aarch64

# 解压安装
# -C 表示切换到某个目录后再解压
# 注意权限问题
sudo tar -C /usr/local -xzf go1.25.3.linux-arm64.tar.gz

# 使环境变量生效
source ~/.bashrc

# 检验是否安装成功
dog@dog:~/install$ go version
go version go1.25.3 linux/arm64
```

开始安装 `gvm`
```shell
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# 检测是否成功
dog@dog:~/install$ gvm version
Go Version Manager v1.0.22 installed at /home/dog/.gvm
```
根据需要可以使用 `gvm` 安装需要的 `Go` 特定版本环境！

