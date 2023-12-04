

简体中文(Simple Chinese)|<a herf="./README-en.md">English</a>

# go_tun2socks_for_windows

## 简介

go_tun2socks基于tun2socks开发的一款游戏加速器的go语言客户端，使用了sqlite3作为数据库存储规则文件。兼容sstap规则文件，当前版本仅支持windows平台。注意当前仅支持`tun2socks`核心。请提交issue提交您希望核心,开发者会添加到`feature`中。

## 使用方法

将tun2socks编译文件放入`core/tun2socks`目录下，规则文件放入`rules`目录下

## 存储原理

第一次使用`rules`使用路由规则把`sstap-rule`规则文件遍历写入`rule.db`数据中,通过数据库查询，达到高速使用路由表。

## 友情链接



## 感谢名单

感谢github给了一个良好的开源平台社区环境
