package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"main/database"
	"os"
	"os/exec"
	"path/filepath"
)

func Runmainwindows() {
	myApp := app.New()
	myWindow := myApp.NewWindow("gotun2socks")
	myWindow.Resize(fyne.NewSize(500, 500))

	label1 := widget.NewLabel("CoreSelect:")
	cb1 := widget.NewSelect([]string{"sing-box", "tun2socks"}, nil)

	label2 := widget.NewLabel("SelectGameRules:")
	cb2 := widget.NewSelect(getFilenamesModel(), nil)
	label3 := widget.NewLabel("ProtocolType:")
	cb3 := widget.NewSelect([]string{"socks", "hysteria2"}, nil)
	button := widget.NewButton("Run", func() {
		if cb1.Selected == "sing-box" && cb3.Selected == "socks" {
			cmd := exec.Command("core/sing-box/sing-box.exe", "run")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// 执行命令
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		} //当选择的协议为socks时，执行sing-box内核的命令
	})

	content := container.NewVBox(
		container.NewHBox(label1, cb1),
		container.NewHBox(label2, cb2),
		container.NewHBox(label3, cb3),
		button,
	)
	iconpath, _ := os.ReadFile("./img/favicon.ico")

	myWindow.SetIcon(fyne.NewStaticResource("icon", iconpath))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	return
}

func getDirectoriesModel(path string) []string {
	var directories []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("获取目录出错", err)
			return err
		}

		if info.IsDir() && filepath.Base(path) != "core" {
			directories = append(directories, path)
		}

		return nil
	})

	if err != nil {
		log.Println("获取目录出错", err)
		return nil
	}

	return directories
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
