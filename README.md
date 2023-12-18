[简体中文(Simplified Chinese)](./README.md)|[English](./README-en.md)|[日本語(Japanese)](./README-jp.md)

# go_tun2socks_for_windows

## 简介

go_tun2socks基于tun2socks与sing-box开发的一款双内核游戏加速器的go语言客户端，使用了sqlite3作为数据库存储规则文件。兼容sstap规则文件，当前版本仅支持windows平台。注意当前仅支持`sing-box`核心。请提交issue提交您希望核心,开发者会添加到`feature`中。

## 使用方法
### 源码安装：
将tun2socks编译文件放入`core/tun2socks`目录下，规则文件放入`rules`目录下;
将sing-box编译文件放入`core/sing-box`目录下，规则文件放入`rules`目录下
### 使用必看:
该项目主要基于tun模式实现，所以需要管理员权限运行，否则无法创建虚拟网卡。如果您不想使用管理员权限运行，可以使用`tap`模式，但是`tap`模式需要您自己安装虚拟网卡驱动，这里不做介绍。如果您不知道如何安装虚拟网卡驱动，建议您使用管理员权限运行。


## 存储原理

第一次使用`rules`使用路由规则把`sstap-rule`规则文件遍历写入`rule.db`数据中,通过数据库查询，达到高速使用路由表。
由于使用了sing-box内核，所以不需要使用`iptables`进行路由转发，所以不需要root权限。第一次使用会自动下载geoip.dat和geosite.dat文件，如果下载失败，可以手动下载，放入项目根目录下。

## 友情链接

## 项目文档

[go_tun2socks_for_windows项目文档](https://github.com/theshdowaura/go_tun2socks_for_windows/wiki)


### 加入我们的群组:
[Telegram群组](https://t.me/gotun2socks/1)
## 感谢名单

感谢github给了一个良好的开源平台社区环境
