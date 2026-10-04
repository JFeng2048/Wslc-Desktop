package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/services"
)

//go:embed all:ui/dist
var assets embed.FS

func main() {
	// 所有 wslc 命令统一通过 executor 执行
	exec := executor.New("wslc")

	systemSvc := services.NewSystemService(exec)
	containerSvc := services.NewContainerService(exec)
	imageSvc := services.NewImageService(exec)
	volumeSvc := services.NewVolumeService(exec)
	networkSvc := services.NewNetworkService(exec)
	commandSvc := services.NewCommandService(exec)

	app := application.New(application.Options{
		Name:        "Wslc Desktop",
		Description: "WSL container desktop manager",
		Services: []application.Service{
			application.NewService(systemSvc),
			application.NewService(containerSvc),
			application.NewService(imageSvc),
			application.NewService(volumeSvc),
			application.NewService(networkSvc),
			application.NewService(commandSvc),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:    "Wslc Desktop",
		Width:    1180,
		Height:   760,
		MinWidth: 960,
		MinHeight: 640,
		// 与前端暗色主题一致的底色，避免加载瞬间白屏
		BackgroundColour: application.NewRGB(14, 17, 22),
		URL:              "/",
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
