package gui

import (
	"testing"

	"github.com/AllenDang/giu"
)

func TeststartEguiWindow(t *testing.T) {
	// 初始化egui窗口
	wnd := giu.NewMasterWindow("tun2socks", 400, 200, 0)
	wnd.Run(updateEguiWindow)
}

func updateEguiWindow() {
	// 处理egui窗口的更新和渲染
	giu.SingleWindow().Layout(
		giu.Label("Hello world from giu"),
	)
}
