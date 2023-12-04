package main

import (
	"go_tun2socks_for_windows/gui"
	"log"
)

func main() {
	if err := gui.RunMainWindow(); err != nil {
		log.Println("启动出错", err)
	}
}
