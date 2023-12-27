package gui

import (
	"fmt"

	"log"
	"os"
	"path/filepath"

	"github.com/getlantern/systray"
)

func OnReady() {
	// 创建一个菜单项3
	menuitemMain := systray.AddMenuItem("显示", "显示界面")
	menuItem := systray.AddMenuItem("退出程序", "退出应用程序")
	// 监听菜单项的点击事件
	go func() {
		<-menuItem.ClickedCh
		myWindow.Close()
		systray.Quit()
	}()
	go func() {
		<-menuitemMain.ClickedCh
		myWindow.Show()

	}()
	// 显示系统托盘图标
	systray.SetTooltip("tun2socks")
	systray.SetTitle("go_tun2socks_for_windows")
	systray.SetIcon(getIconData())

	// 在后台运行
	fmt.Println("应用程序已启动，可以在系统托盘中找到图标")
	return
}

func OnExit() {
	// 清理资源、关闭程序等操作
	fmt.Println("应用程序已退出")
	os.Exit(0)

}

func getIconData() []byte {
	// 从文件中读取图标数据
	imgPath, err := filepath.Abs("gui/img/favicon.ico")
	img, err := os.ReadFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
