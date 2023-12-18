[简体中文(Simplified Chinese)](./README.md)|[English](./README-en.md)|[日本語](./README-jp.md)

# go_tun2socks_for_windows

## Introduction

go_tun2socks_for_windows is a Go language client for a dual-kernel game accelerator based on tun2socks and sing-box. It uses sqlite3 as the database to store rule files. It is compatible with sstap rule files and currently supports only the Windows platform. Note that currently only the `sing-box` kernel is supported. Please submit an issue if you wish to add a different kernel, and the developer will add it to the `feature` list.

## Usage
### Source code installation:
Place the compiled tun2socks file in the `core/tun2socks` directory, and place the rule file in the `rules` directory.
Place the compiled sing-box file in the `core/sing-box` directory, and place the rule file in the `rules` directory.
### Important Note:
This project is mainly based on the tun mode, so it needs to be run with administrator privileges; otherwise, it will not be able to create a virtual network card. If you do not want to run with administrator privileges, you can use the `tap` mode, but the `tap` mode requires you to install the virtual network card driver yourself, which is not covered here. If you do not know how to install the virtual network card driver, it is recommended to run with administrator privileges.

## Storage Principle

The first time you use the `rules`, the routing rules from the `sstap-rule` rule file are traversed and written into the `rule.db` database. By querying the database, high-speed routing table usage is achieved.
Since the sing-box kernel is used, there is no need to use `iptables` for routing forwarding, so root permission is not required. The geoip.dat and geosite.dat files will be automatically downloaded the first time you use it. If the download fails, you can manually download them and place them in the project root directory.

## Links

## Project Documentation

[go_tun2socks_for_windows Project Documentation](https://github.com/theshdowaura/go_tun2socks_for_windows/wiki)


### Join our group:
[Telegram Group](https://t.me/gotun2socks/1)

## Acknowledgements

Thanks to GitHub for providing a good open-source platform community environment.