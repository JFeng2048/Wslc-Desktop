# AGENT.md

Wslc Desktop —— 基于 Wails v3 的 WSL 容器图形化管理工具（Windows only）。

## 技术栈

- **后端**：Go 1.26+，Wails v3（`github.com/wailsapp/wails/v3`）。
- **前端**：Vue 3 + TypeScript + Vite + Pinia + vue-i18n + Naive UI + lucide-vue-next。
- **命令层**：所有 WSL 容器操作统一通过调用外部 CLI `wslc` 完成（不直连 Docker/WSL API）。
  `wslc` 是微软的 WSL 容器 CLI，支持 34 个子命令。

## 目录结构

```
main.go                      入口：构造 executor 与 6 个 service，注册到 Wails
internal/
  executor/executor.go       统一封装 wslc（及 wsl 等）命令调用，返回 Result{Stdout,Stderr,ExitCode}
  models/models.go           与 wslc JSON 输出对应的数据模型（camelCase json 标签）
  services/
    decode.go                通用 JSON 解析器：大小写不敏感，兼容 wslc 的 PascalCase 键与 JSONL/数组
    container.go             容器：list/start/stop/restart/kill/rm/prune/inspect/logs/stats
    image.go                 镜像：list/rmi/inspect/pull/push/tag/prune
    volume.go                卷：list/rm/prune/inspect/create
    network.go               网络：list/rm/prune/inspect/create
    system.go                系统：GetStatus/Info/Version
    command.go               CommandService.Run：运行任意 wslc 子命令（命令终端用）
build/                       Wails 构建配置（仅保留 windows 平台）
ui/
  src/
    bindings/                wails 生成的 TS 绑定（勿手改，由 `wails3 generate bindings` 产出）
    api/                     对绑定的薄封装（按资源分文件）
    stores/                  Pinia 状态（每个资源一个 store）
    views/                   页面：ContainerList / ImageList / VolumeList / NetworkList / SystemInfo / Terminal
    components/JsonDrawer.vue 复用的详情抽屉（展示 inspect/logs/stats 原始输出）
    layouts/MainLayout.vue   侧边栏导航 + 顶栏（主题/语言切换）
    router/ i18n/ types/     路由、中英文文案、与后端对齐的类型
```

## wslc 命令覆盖

| 资源 | 结构化页面操作 | 底层 wslc 子命令 |
|------|----------------|------------------|
| 容器 | 列表/启动/停止/重启/强杀/删除/清理/详情/日志/资源 | `list -a --format json`, `start`, `stop`, `restart`, `kill`, `rm -f`, `container prune`, `inspect`, `logs`, `stats --format json` |
| 镜像 | 列表/删除/详情/拉取/推送/打标签/清理 | `images --format json`, `rmi -f`, `inspect`, `pull`, `push`, `tag`, `image prune` |
| 卷   | 列表/删除/清理/详情/创建 | `volume ls --format json`, `volume rm`, `volume prune`, `inspect`, `create` |
| 网络 | 列表/删除/清理/详情/创建 | `network ls --format json`, `network rm`, `network prune`, `inspect`, `create` |
| 系统 | 状态/信息/版本 | `version`, `info --format json` |
| 其余命令 | 命令终端（Terminal 页）覆盖全部 34 个子命令 | `attach/build/create/exec/events/export/import/load/login/logout/run/save/tag/...` |

> 注意：`wslc list`/`images`/`volume ls`/`network ls` 的 `--format` 仅接受 `json` 或 `table`，
> **不支持 Go 模板**，因此列表一律使用 `--format json` 并经由 `decode.go` 解析。

## 本地开发

前置：Go、Node/npm、已安装并在 PATH 中的 `wslc` 与 `wsl`。

```powershell
# 启动开发（前端热更新 + 桌面窗口）
wails3 dev

# 仅重新生成前端绑定（改了 internal/services 或 models 后必须执行）
wails3 generate bindings -clean=true -ts -i -names -d ui/src/bindings

# 生产构建
wails3 build
```

- 绑定生成命令已固化在 `build/Taskfile.yml` 的 `generate:bindings` 任务中，
  `wails3 build` / `wails3 dev` 会自动调用。
- 本项目仅面向 Windows；`build/` 下只保留 `windows` 平台，其他平台目录为占位空壳
  （wails 加载阶段会枚举各平台 Taskfile，缺失会导致构建报错）。

## 约定

- 修改后端 service / model 后，必须重新生成绑定并同步 `ui/src/types/index.ts`（已改为从绑定再导出）。
- 前端类型以 `ui/src/bindings/...` 为唯一来源，`ui/src/types/index.ts` 仅做再导出，避免漂移。
- 列表解析一律走 `--format json` + `decodeJSONLines`，新增资源类型时复用该函数即可。
- `CommandService.Run` 统一加 60s 超时，命令终端不支持 `-f` 等持续跟随命令，以免阻塞 UI。
