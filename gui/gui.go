package gui

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"go_tun2socks_for_windows/database"
	"log"
	"os"
	"path/filepath"
)

func RunMainWindow() error {
	var (
		label1 *walk.Label
		cb1    *walk.ComboBox

		label2 *walk.Label
		cb2    *walk.ComboBox
	)

	mw := declarative.MainWindow{
		Title:  "gotun2socks",
		Size:   declarative.Size{Width: 500, Height: 500},
		Layout: declarative.VBox{MarginsZero: true},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.Label{
						AssignTo: &label1,
						Text:     "内核选择:",
					},
					declarative.ComboBox{
						AssignTo: &cb1,
						Model:    getDirectoriesModel("core"),
					},
					declarative.Label{
						AssignTo: &label2,
						Text:     "游戏规则选择:",
					},
					declarative.ComboBox{
						AssignTo: &cb2,
						Model:    getFilenamesModel(),
					},
				},
			},

			declarative.PushButton{
				Text: "确定",
				OnClicked: func() {
					database.DbOut(cb2.Text())
				},
			},
		},
	}

	_, err := mw.Run()
	return err
}

func getDirectoriesModel(path string) []string {
	var directories []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("获取目录出错", err)
			return nil
		}

		if info.IsDir() && path != "core" {
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
	return names
}
