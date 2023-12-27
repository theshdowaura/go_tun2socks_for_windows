package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/getlantern/systray"
	"log"
	"main/config"
	"main/database"
	"os"
	"os/exec"
)

var counter int
var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("gotun2socks")
var cb1 *widget.Select = widget.NewSelect([]string{"sing-box", "tun2socks"}, nil)
var cb2 *widget.Select = widget.NewSelect(getFilenamesModel(), nil)
var cb3 *widget.Select = widget.NewSelect([]string{"socks", "hysteria2"}, nil)

func Runmainwindows() {
	myWindow.Resize(fyne.NewSize(500, 500))

	label1 := widget.NewLabel("CoreSelect:")

	label2 := widget.NewLabel("SelectGameRules:")
	label3 := widget.NewLabel("ProtocolType:")
	button1 := widget.NewButton("Run", func() {
		if cb1.Selected == "sing-box" && contains(getFilenamesModel(), cb2.Selected) && cb3.Selected == "socks" {
			myWindow.Hide()
			NewConfigRoute(database.DbOut(cb2.Selected)) //更新路由地址
			cmd := exec.Command("core/sing-box/sing-box.exe", "run")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			// 执行命令
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)

			}

		} //当选择的协议为socks时，执行sing-box内核的命令
	})
	button2 := widget.NewButton("Hide", func() {
		myWindow.Hide()
		systray.Run(OnReady, nil)
	})

	content := container.NewVBox(
		container.NewHBox(label1, cb1),
		container.NewHBox(label2, cb2),
		container.NewHBox(label3, cb3),
		button1, button2,
	)
	iconpath, _ := os.ReadFile("./img/favicon.ico")
	myWindow.SetIcon(fyne.NewStaticResource("icon", iconpath))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	systray.Run(OnReady, nil)
	return
}

func contains(s []string, searchterm string) bool {
	for _, v := range s {
		if v == searchterm {
			return true
		}
	}
	return false
}

func getFilenamesModel() []string {
	names, err := database.GetDistinctFilenames()
	if err != nil {
		log.Println("获取文件名出错", err)
		return nil
	}
	if len(names) == 0 {
		log.Println("No distinct filenames found")
		return nil
	}
	return names
}

func NewConfigRoute(ips []string) {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 修改inet4_route_address字段的值
	cfg.Inbounds[0].Inet4RouteAddress = ips

	// 将修改后的配置文件保存到新文件
	err = config.SaveConfig(cfg, "config.json")
	if err != nil {
		log.Fatalf("保存配置文件失败: %v", err)
	}

	fmt.Println("配置文件修改成功！")
}
func showAndSelect() {
	myWindow.Show()
	// 等待选择完成
	for {
		if cb1.Selected == "sing-box" && contains(getFilenamesModel(), cb2.Selected) && cb3.Selected == "socks" {
			Runmainwindows() // 重新运行主窗口函数
			break
		} else {
			// 等待选择完成
		}
	}
}
