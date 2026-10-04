<div align="center">

# Wslc Desktop

**面向 Windows 的 WSL 容器图形化桌面管理器 —— `wslc` CLI 的原生图形界面。**

[![Platform](https://img.shields.io/badge/platform-Windows-0078D4?logo=windows&logoColor=white)](https://learn.microsoft.com/windows/wsl/)
[![Wails](https://img.shields.io/badge/Wails-v3-06B6D4?logo=go&logoColor=white)](https://v3.wails.io/)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
![License](https://img.shields.io/badge/license-unlicensed-lightgrey)

[English](../README.md) | **简体中文**

</div>

---

Wslc Desktop 是一个通过图形界面管理 WSL 容器的 Windows 桌面应用，无需记忆命令。它封装了微软官方的
[`wslc`](https://learn.microsoft.com/windows/wsl/) CLI，因此容器、镜像、卷与网络仍由你在终端里
使用的同一套工具处理，只是把命令换成了表格、按钮与表单。对于 WSL 场景，它可作为 Docker Desktop
的轻量替代方案，支持 Windows 10 与 Windows 11（需启用 WSL 2）。

基于 Wails v3、Go 1.25、Vue 3 与 TypeScript 构建。

## 快速开始

只需两个命令。先准备环境：Windows 10/11 并启用 WSL 2、`wslc` CLI 已在 `PATH` 中、
[Go](https://go.dev/dl/) 1.25+、[Node.js](https://nodejs.org/) 20+，以及
[Wails v3 CLI](https://v3.wails.io/getting-started/installation/)。

```bash
git clone https://github.com/JFeng2048/Wslc-Desktop.git
cd Wslc-Desktop
cd ui && npm install && cd ..

wails3 dev     # 启动开发模式（热重载）
wails3 build   # 生产构建 -> bin/wslc-desktop-amd64.exe
```

`wails3 dev` 会同时启动 Vite 开发服务器与桌面窗口，保存 Go 或 TypeScript 改动即可热更新。
`wails3 build` 产出免安装、零依赖的 Windows 可执行文件。

## 功能特性

- **系统状态** – WSL 可用性与版本、`wslc` 守护进程状态与版本
- **容器** – 列表、搜索、按状态筛选、启动、停止、重启、强杀、删除、日志、资源占用、清理
- **镜像** – 列表、搜索、拉取、推送、重新打标签、删除、清理
- **卷** – 列表、搜索、创建、删除、清理
- **网络** – 列表、搜索、创建、删除、清理
- **命令终端** – 执行任意 `wslc` 子命令，查看 stdout、stderr 与退出码
- **便携 ZIP 发布** – 单一 `.exe` 即可分发，无需安装器与额外运行时
- **中英文界面**，浅色主题

## 应用截图

| 卷 | 镜像 | 网络 |
| --- | --- | --- |
| ![卷列表](images/docs-volumes.png) | ![镜像列表](images/docs-image.png) | ![网络列表](images/docs-network.png) |

## 构建与打包

默认目标架构为 `amd64`。传入 `ARCH` 即可交叉编译，架构始终体现在可执行文件名中，
因此多种架构的产物可以共存于 `bin/`。

```bash
wails3 task build ARCH=arm64              # -> bin/wslc-desktop-arm64.exe
wails3 task build:all                     # 同时构建 amd64 与 arm64
```

ZIP 压缩包内部的可执行文件统一命名为 `wslc-desktop.exe`，版本与架构体现在压缩包文件名上：

```bash
wails3 task package:zip                   # -> bin/wslc-desktop-0.1.0-amd64.zip
wails3 task package:zip:all               # 同时生成两种架构的 ZIP
wails3 task package:zip APP_VERSION=0.2.0 # 覆盖文件名中的版本号
```

| `bin/` 产物 | 说明 |
| --- | --- |
| `wslc-desktop-amd64.exe` | amd64 可执行文件 |
| `wslc-desktop-arm64.exe` | arm64 可执行文件 |
| `wslc-desktop-0.1.0-amd64.zip` | 便携版 ZIP，内部为 `wslc-desktop.exe`（amd64） |
| `wslc-desktop-0.1.0-arm64.zip` | 便携版 ZIP，内部为 `wslc-desktop.exe`（arm64） |

ZIP 文件名中的版本号取自 `Taskfile.yml` 的 `APP_VERSION`，请与 `build/config.yml` 中的
`info.version` 保持一致。

## 目录结构

```text
main.go                入口：装配 executor 与各 service 并注册到 Wails
internal/executor/     基于 wslc / wsl 的命令执行层
internal/models/       与 wslc JSON 输出对应的数据模型
internal/services/     业务逻辑，按资源分文件
build/                 Wails 构建配置（仅 Windows）与图标资源
ui/src/bindings/       自动生成的 TS 绑定（已 Git 忽略）
ui/src/api/            对绑定层的薄封装
ui/src/stores/         Pinia store，按资源划分
ui/src/views/          系统、容器、镜像、卷、网络、命令终端
docs/                  补充文档
```

## 常见问题

**需要安装 Docker Desktop 吗？**
不需要。本应用驱动的是 WSL 2 自带的 `wslc` CLI，不依赖 Docker Desktop。

**最低环境要求是什么？**
Windows 10 build 1809 及以上（推荐 Windows 11）、WSL 2、`wslc` CLI 位于 `PATH`、
Go 1.25+、Node.js 20+，开发还需 Wails v3 CLI。

**支持 Windows on ARM 吗？**
支持。使用 `wails3 task build ARCH=arm64` 构建，或用 `build:all` / `package:zip:all`
一次性产出 amd64 与 ARM64 两种产物。

**没有独立页面的 `wslc` 命令怎么执行？**
在命令终端页可执行全部 34 个 `wslc` 子命令，并查看 stdout、stderr 与退出码。

**应用是否自带运行时？**
不自带。可执行文件仅依赖 Windows 10/11 已有的 WebView2 运行时，ZIP 解压后可直接复制到其他机器运行。

**可以切换界面语言吗？**
可以，点击顶栏的语言按钮即可在运行时切换中英文。

## 文档

- [English README](../README.md)
- [AGENT.md](../AGENT.md) – 架构说明与开发约定，修改本代码库时值得一读
- [Wails v3 文档](https://v3.wails.io/)
- [WSL 官方文档](https://learn.microsoft.com/windows/wsl/)

## 参与贡献

- 发现 Bug？[提交 Issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=bug)
- 需要新功能？[提交 Issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=enhancement)
- 想贡献代码？[Fork 本仓库](https://github.com/JFeng2048/Wslc-Desktop/fork)，创建分支后
  [发起 Pull Request](https://github.com/JFeng2048/Wslc-Desktop/pulls/new)
- 有疑问？[发起讨论](https://github.com/JFeng2048/Wslc-Desktop/discussions)

## 许可协议

本项目尚未声明开源许可证。在许可证确定之前，著作权归作者所有。若你计划再分发本项目代码或基于其
二次开发，请先[提交 Issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new)沟通。

## 作者

**JFeng2048** — [GitHub](https://github.com/JFeng2048) · [JFeng2048@outlook.com](mailto:JFeng2048@outlook.com)
