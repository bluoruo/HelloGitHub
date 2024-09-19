# Hello GitHub
简单的用golang写的自动查找最新的 github.com IP地址。 \
解决不能访问github.com，这个程序必须运行在没有dns污染的环境<sup>(1)</sup>，请求的dns服务器必须是未有地址规则的<sup>(2)</sup>。\
本程序支持 windows、Linux、FreeBSD、MacOS和Android等设备。

### 编译
golang版本 大于 1.19.6
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

#### Windows 下编译所有版本

```bash
# windows 版本 直接编译是你当前系统能使用的版本
go build -ldflags "-w -s" .

# Linux - x86_64 版本
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" .

# Linux - arm64 版本
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -ldflags "-w -s" .

# FreeBSD - x86_64 版本
SET CGO_ENABLED=0
SET GOOS=freebsd
SET GOARCH=amd64
go build -ldflags "-w -s" .

# MacOs - Intel 版本
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" .


# MacOs - M1/M2 版本
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=arm64
go build -ldflags "-w -s" .
```
#### Linux 下编译所有版本
```bash
# Linux 版本 直接编译是你当前系统能使用的版本
go build -ldflags "-w -s" .

# Windows - x86_64 版本
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" .

# Windows - arm64 版本
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "-w -s" .

# FreeBSD - x86_64 版本
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags "-w -s" .

# MacOs - Intel 版本
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" .

# MacOs - M1/M2 版本
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s" .
```
#### FreeBSD 和 MacOS 下编译和 Linux 一样。

###### 特别说明：编译 Android 版本，需要先下载 Android SDK 编译时候改 GoOS=Android,GoArch 为对应的CPU版本就可以了。

如果不想编译可以直接下载编译好的版本。

### 使用
+ 下载好程序，首先要在程序根目录添加 dns.txt 是dns服务器列表，每行一个。比如 `218.85.152.99`(当然这个不能用)。
+ + 这里不能使用有规则的dns服务器比如国内的`114.114.114.114` 国外的如：`8.8.8.8 1.1.1.1` 之类的。
+ + domain.txt 放置的是要解析的域名列表，每行一个。我们给出的 github用到的和vscode的域名，你可以自己添加和修改这些。
+ Windows 可以加入计划任务 每天那个时间点运行
+ Linux 可以使用 cron 计划任务来设置每天运行
+ 执行完程序根目录回放置一个hosts文件，直接使用即可，也可以根据你的需要在 temp目录找你觉得适合的。
+ 电信推荐使用US开头的，联通适合选择DE开头的。

#### 你可以直接下载提供的hosts文件

##### Windows
修改 `C:\windows\system32\drivers\etc\hosts` \
如果没有这个文件可以创建.
##### Linux/FreeBSD/OpenWrt/MacOs
修改 `/etc/hosts`

##### 可以直接下面的文件到你的hosts文件

```bash

# HelloGitHub Host Start

140.82.113.26            alive.github.com
140.82.121.6             api.github.com
185.199.109.153          assets-cdn.github.com
185.199.111.133          avatars.githubusercontent.com
185.199.111.133          avatars0.githubusercontent.com
185.199.109.133          avatars1.githubusercontent.com
185.199.110.133          avatars2.githubusercontent.com
185.199.111.133          avatars3.githubusercontent.com
185.199.108.133          avatars4.githubusercontent.com
185.199.108.133          avatars5.githubusercontent.com
185.199.111.133          camo.githubusercontent.com
140.82.113.22            central.github.com
185.199.110.133          cloud.githubusercontent.com
140.82.121.9             codeload.github.com
140.82.113.22            collector.github.com
185.199.110.133          desktop.githubusercontent.com
185.199.108.133          favicons.githubusercontent.com
140.82.121.4             gist.github.com
3.5.30.33                github-cloud.s3.amazonaws.com
54.231.226.217           github-com.s3.amazonaws.com
16.182.109.17            github-production-release-asset-2e65be.s3.amazonaws.com
3.5.28.21                github-production-repository-file-5c1aeb.s3.amazonaws.com
3.5.25.221               github-production-user-asset-6210df.s3.amazonaws.com
192.0.66.2               github.blog
140.82.121.4             github.com
140.82.112.17            github.community
185.199.111.154          github.githubassets.com
146.75.117.194           github.global.ssl.fastly.net
185.199.109.153          github.io
185.199.111.133          github.map.fastly.net
185.199.109.153          githubstatus.com
140.82.112.26            live.github.com
185.199.111.133          media.githubusercontent.com
185.199.108.133          objects.githubusercontent.com
13.107.42.16             pipelines.actions.githubusercontent.com
185.199.111.133          raw.githubusercontent.com
185.199.111.133          user-images.githubusercontent.com
13.107.246.45            vscode.dev


# Update time: 2024-09-08 05:32:15
# Update url: https://raw.githubusercontent.com/bluoruo/HelloGitHub/master/hosts
# Star me: https://github.com/bluoruo/HelloGitHub
# HelloGitHub Host End

```

##### 修改完成后重启你的浏览器就可以生效了

~~ 本文件每天更新
### 服务器到期暂停更新

### 注脚解释
(1) DNS污染存在于大部分地方，检查是否被DNS污染可以使用tcpDNS协议。\
(2) 规则判断的DNS服务器，会根据你的请求IP位置返回指定的解析地址。

--------------
### 开源声明
HelloGitHub By Comanche Lab.  基于GPL V3 协议开源。
































































































































































































































































































































































































































































































































































































