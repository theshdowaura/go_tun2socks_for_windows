[TOC]

[简体中文(Simplified Chinese)](./README.md)|[English](./README-en.md)

# go_tun2socks_for_windows

## Introduction

go_tun2socks is a Go language client for a game accelerator based on tun2socks. It uses SQLite3 as the database to store rule files. It is compatible with sstap rule files and currently only supports the Windows platform. Please submit an issue if you want to add support for other cores, and the developer will add it to the features.

## Usage

Place the compiled tun2socks file in the `core/tun2socks` directory and the rule file in the `rules` directory.

## Storage Principle

The first time you use the `rules`, the routing rules in the `sstap-rule` rule file will be traversed and written into the `rule.db` database. The routing table is then used for fast lookup through database queries.

## Links

## Acknowledgments

Thanks to GitHub for providing a good open-source platform and community environment.