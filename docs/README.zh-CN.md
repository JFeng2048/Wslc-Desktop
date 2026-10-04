<div align="center">

# Wslc Desktop

**面向 Windows 的 WSL 容器图形化桌面管理器。**

[![Platform](https://img.shields.io/badge/platform-Windows-0078D4?logo=windows&logoColor=white)](https://learn.microsoft.com/windows/wsl/)
[![Wails](https://img.shields.io/badge/Wails-v3-06B6D4?logo=go&logoColor=white)](https://v3.wails.io/)
[![Vue](https://img.shields.io/badge/Vue-3.5-42b883?logo=vuedotjs&logoColor=white)](https://vuejs.org/)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
![License](https://img.shields.io/badge/license-unlicensed-lightgrey)

[English](../README.md) | **简体中文**

</div>

---

## 项目简介

**Wslc Desktop** 把微软官方的 [`wslc`](https://learn.microsoft.com/windows/wsl/) 命令行工具包装成一个
快速、原生的 Windows 桌面应用。你不必再背诵几十个子命令和参数，日常容器运维通过一个简洁、可搜索的
界面即可完成。

所有容器操作都委托给官方 `wslc` CLI 执行，而不直接对接 Docker / WSL 守护进程，因此行为与你在终端里
敲命令完全一致。

## 功能特性

| 模块 | 能力 |
| --- | --- |
| **系统状态** | WSL 可用性与版本、`wslc` 守护进程状态、版本与会话管理器信息 |
| **容器** | 列表、按状态筛选、搜索、启动、停止、重启、强杀、强制删除、清理、详情、日志、实时资源占用 |
| **镜像** | 列表、搜索、拉取、推送、重新打标签、详情、删除、清理 |
| **卷** | 列表、搜索、创建、详情、删除、清理 |
| **网络** | 列表、搜索、创建、详情、删除、清理 |
| **命令终端** | 执行 `wslc` 全部 34 个子命令，展示 stdout / stderr / 退出码 |

补充说明：

- **详情抽屉** – 复用的抽屉组件展示 `inspect`、`logs`、`stats` 原始输出，便于复制。
- **命令终端** – 为没有独立页面的长尾 `wslc` 子命令提供的内置控制台。
- **仅浅色主题** – 只保留一套经过调校的浅色配色，聚焦可读性。
- **中英文双语** – 界面文案完整本地化，可在运行时切换。
- **搜索与筛选** – 所有列表视图均支持关键字搜索，容器列表额外支持状态筛选。
- **自适应表格** – 表格支持横向滚动，操作列固定在右侧。

## 技术栈

| 层次 | 技术 |
| --- | --- |
| 桌面外壳 | [Wails v3](https://v3.wails.io/)（Windows 上使用 WebView2） |
| 后端 | Go 1.25 |
| 前端 | Vue 3 · TypeScript · Vite |
| 状态管理 | [Pinia](https://pinia.vuejs.org/zh/) |
| 路由 | [vue-router](https://router.vuejs.org/zh/) |
| UI 组件库 | [Naive UI](https://www.naiveui.com/) + [Lucide](https://lucide.dev/) 图标 |
| 国际化 | [vue-i18n](https://vue-i18n.intlify.dev/zh/) |
| 任务运行器 | [go-task](https://taskfile.dev/) |
| 打包 | NSIS 安装包与 MSIX 包 |

## 环境要求

- Windows 10（1809 及以上）或 Windows 11
- 已启用 WSL 2，且 [`wslc`](https://learn.microsoft.com/windows/wsl/) 容器 CLI 已在 `PATH` 中
- [Go](https://go.dev/dl/) 1.25 或更高版本
- [Node.js](https://nodejs.org/) 20 或更高版本（自带 npm）

以下为可选项，仅在使用相关工具与打包命令时需要：

- [Wails v3 CLI](https://v3.wails.io/getting-started/installation/)：
  `go install github.com/wailsapp/wails/v3/cmd/wails3@latest`
- [go-task](https://taskfile.dev/docs/installation)：
  `go install github.com/go-task/task/v3/cmd/task@latest`
- [NSIS](http://nsis.sourceforge.net/Main_Page)——仅打包 NSIS 安装包时需要
- MSIX 打包还需执行 `task install:msix:tools` 安装 MSIX 工具链

## 快速开始

```bash
# 1. 克隆仓库
git clone https://github.com/JFeng2048/Wslc-Desktop.git
cd Wslc-Desktop

# 2. 安装前端依赖
cd ui && npm install && cd ..

# 3. 以开发模式启动（Go 与前端均支持热更新）
wails3 dev
```

Vite 开发服务器与桌面窗口会自动启动，保存 Go 或 TypeScript 改动即可即时生效。

> `ui/src/bindings/` 下的 TypeScript 绑定由 `wails3 generate bindings` 生成，
> `wails3 dev` 与 `wails3 build` 会自动执行该步骤。该目录已被 Git 忽略，切勿手动修改。

## 常用命令

以下均为 [go-task](https://taskfile.dev/) 任务，括号中给出等价的 `wails3` 命令。

| 命令 | 等价命令 | 说明 |
| --- | --- | --- |
| `task dev` | `wails3 dev` | 启动开发模式（热重载） |
| `task build` | `wails3 build` | 构建 `bin/wslc-desktop.exe` |
| `task run` | — | 运行已构建的可执行文件 |
| `task package` | — | 构建 NSIS 安装包（默认） |
| `task install:msix:tools` | — | 安装 MSIX 打包工具链 |

常用变体：

```bash
task package INSTALL_SCOPE=user            # 按用户安装的 NSIS 安装包
task package FORMAT=msix                   # MSIX 包（发布者 CN=JFeng2048）
task build DEV=true                        # 开发模式构建
```

## 构建与打包

```bash
# 生产构建 -> bin/wslc-desktop.exe
task build

# NSIS 安装包
task package

# MSIX 包
task package FORMAT=msix
```

应用图标由 `build/logo.png` 生成：

```bash
# 依据源 logo 重新生成 build/windows/icon.ico
wails3 generate icons -input build/logo.png -windowsfilename build/windows/icon.ico
```

代码签名为可选项。可通过 Wails CLI（`wails3 setup signing`）配置证书，或在调用
`sign`、`sign:installer` 任务时设置 `SIGN_CERTIFICATE` / `SIGN_THUMBPRINT`。

## 目录结构

```text
main.go                    入口：装配 executor 与各 service 并注册到 Wails
internal/
  executor/                基于 wslc / wsl 的命令执行层
  models/                  与 wslc JSON 输出对应的数据模型
  services/                业务逻辑，按资源分文件
build/                     Wails 构建配置（仅 Windows）与图标资源
ui/
  src/
    bindings/              自动生成的 TS 绑定（已 Git 忽略）
    api/                   对绑定层的薄封装
    stores/                Pinia store，按资源划分
    views/                 系统、容器、镜像、卷、网络、命令终端
    layouts/               侧边导航与顶栏
    i18n/                  中英文文案
```

## 架构说明

- 所有列表均以 `--format json` 拉取，再经由共用的大小写不敏感解析器处理，因此 `wslc` 的
  PascalCase 字段与 JSON Lines 输出都能正确解析。
- `wslc` 的 `--format` 不支持 Go 模板，因此统一使用 JSON。
- 命令终端中执行的任意命令均有 60 秒超时，并有意不暴露 `-f` 等持续跟随类参数，确保界面不会被阻塞。
- 前端类型从生成的绑定再导出，保证 Go 与 TypeScript 模型自动保持一致。

## 参与贡献

欢迎贡献代码。如果是较大的改动，建议先提 Issue 对齐方向，避免投入无用功。

- **发现 Bug？** [提交 Issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=bug)
- **需要新功能？** [提交 Issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=enhancement)
- **想贡献代码？**
  1. [Fork 本仓库](https://github.com/JFeng2048/Wslc-Desktop/fork)
  2. 创建主题分支：`git checkout -b feat/my-change`
  3. 进行清晰的原子化提交——中文或英文提交信息均可
  4. 推送分支到你的 Fork
  5. [创建 Pull Request](https://github.com/JFeng2048/Wslc-Desktop/pulls/new)
- **疑问或想法？** [发起讨论](https://github.com/JFeng2048/Wslc-Desktop/discussions)

如果你修改了 Go 后端，提交前请记得重新生成绑定：

```bash
wails3 generate bindings -clean=true -ts -i -names -d ui/src/bindings
```

## 问题反馈与支持

- Bug 报告与功能需求：[Issues](https://github.com/JFeng2048/Wslc-Desktop/issues)
- 疑问：[Discussions](https://github.com/JFeng2048/Wslc-Desktop/discussions)
- `wslc` 本身的问题：[WSL 官方文档](https://learn.microsoft.com/windows/wsl/)
- 框架相关问题：[Wails v3 文档](https://v3.wails.io/) 与
  [Wails 社区](https://github.com/wailsapp/wails/discussions)

## 许可协议

本项目尚未声明开源许可证。在许可证确定之前，著作权归作者所有。若你计划再分发本项目代码或基于其
二次开发，请先[提交 Issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new)沟通。

## 致谢

- [Wails](https://wails.io/)——应用框架
- 微软 [`wslc`](https://learn.microsoft.com/windows/wsl/)——被管理的容器 CLI
- [Naive UI](https://www.naiveui.com/)、[Lucide](https://lucide.dev/)、[Pinia](https://pinia.vuejs.org/)、
  [vue-i18n](https://vue-i18n.intlify.dev/)——UI、图标、状态与国际化库
- [go-task](https://taskfile.dev/)——构建流水线使用的任务运行器

## 作者

**JFeng2048** — [GitHub](https://github.com/JFeng2048) · [JFeng2048@outlook.com](mailto:JFeng2048@outlook.com)
