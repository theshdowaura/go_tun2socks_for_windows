package gui

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/getlantern/systray"
)

var (
	databasepath string
)

func OnReady() {
	// 创建一个菜单项
	menuItem := systray.AddMenuItem("退出", "退出应用程序")
	menuItem_main := systray.AddMenuItem("开始", "显示界面")
	databasepath = "../database/rules.db"
	// 监听菜单项的点击事件
	go func() {
		<-menuItem.ClickedCh

		systray.Quit()
	}()
	go func() {
		<-menuItem_main.ClickedCh

	}()
	// 显示系统托盘图标
	systray.SetTooltip("tun2socks")
	systray.SetTitle("go_tun2socks_for_windows")
	systray.SetIcon(getIconData())

	// 在后台运行
	fmt.Println("应用程序已启动，可以在系统托盘中找到图标")
}

func OnExit() {
	// 清理资源、关闭程序等操作
	fmt.Println("应用程序已退出")
	os.Exit(0)
}

func getIconData() []byte {
	// 从文件中读取图标数据
	imgPath, err := filepath.Abs("./img/favicon.ico")
	img, err := os.ReadFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func queryFilename(filename string) {
	db, err := sql.Open("sqlite3", databasepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT ips FROM rule WHERE filename = '%s'", filename)
	ipsRow := db.QueryRow(query)
	var ips string
	err = ipsRow.Scan(&ips)
	if err != nil {
		log.Fatal(err)
	}

	// 使用filename和ips作为tun2socks的路由参数
	fmt.Printf("filename: %s, ips: %s\n", filename, ips)
}
