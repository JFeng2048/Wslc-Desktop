<div align="center">

# Wslc Desktop

**A graphical desktop manager for WSL containers on Windows — a native GUI for the `wslc` CLI.**

[![Platform](https://img.shields.io/badge/platform-Windows-0078D4?logo=windows&logoColor=white)](https://learn.microsoft.com/windows/wsl/)
[![Wails](https://img.shields.io/badge/Wails-v3-06B6D4?logo=go&logoColor=white)](https://v3.wails.io/)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
![License](https://img.shields.io/badge/license-unlicensed-lightgrey)

**English** | [简体中文](docs/README.zh-CN.md)

</div>

---

Wslc Desktop is a Windows desktop application for managing WSL containers through a graphical
interface instead of the command line. It wraps Microsoft's official
[`wslc`](https://learn.microsoft.com/windows/wsl/) CLI, so containers, images, volumes and networks
are handled by the same tooling you would use in a terminal — only the commands are replaced by
tables, buttons and forms. It is a lightweight alternative to Docker Desktop for WSL workloads, and
it runs on Windows 10 and Windows 11 with WSL 2.

Built with Wails v3, Go 1.25, Vue 3 and TypeScript.

## Quick start

Two commands are all you need. Install the prerequisites first: Windows 10/11 with WSL 2, the
`wslc` CLI on `PATH`, [Go](https://go.dev/dl/) 1.25+, [Node.js](https://nodejs.org/) 20+, and the
[Wails v3 CLI](https://v3.wails.io/getting-started/installation/).

```bash
git clone https://github.com/JFeng2048/Wslc-Desktop.git
cd Wslc-Desktop
cd ui && npm install && cd ..

wails3 dev     # run the app with hot reload
wails3 build   # production build -> bin/wslc-desktop-amd64.exe
```

`wails3 dev` starts the Vite dev server and the desktop window together; Go and TypeScript changes
reload on save. `wails3 build` produces a standalone, dependency-free Windows executable.

## Features

- **System** – WSL availability and version, `wslc` daemon status and version
- **Containers** – list, search, filter by state, start, stop, restart, kill, remove, logs, stats, prune
- **Images** – list, search, pull, push, retag, remove, prune
- **Volumes** – list, search, create, remove, prune
- **Networks** – list, search, create, remove, prune
- **Terminal** – run any `wslc` subcommand and read stdout, stderr and the exit code
- **Portable ZIP builds** – ship a single `.exe`, no installer or runtime required
- **English and Chinese** interface, light theme

## Screenshots

| Volumes | Images | Networks |
| --- | --- | --- |
| ![Volumes view](docs/images/docs-volumes.png) | ![Images view](docs/images/docs-image.png) | ![Networks view](docs/images/docs-network.png) |

## Building and packaging

The default target is `amd64`. Pass `ARCH` to cross-compile; the architecture is always part of the
binary name, so several builds can coexist in `bin/`.

```bash
wails3 task build ARCH=arm64              # -> bin/wslc-desktop-arm64.exe
wails3 task build:all                     # both architectures
```

ZIP archives contain the executable named `wslc-desktop.exe`, while the archive name carries the
version and architecture:

```bash
wails3 task package:zip                   # -> bin/wslc-desktop-0.1.0-amd64.zip
wails3 task package:zip:all               # ZIPs for both architectures
wails3 task package:zip APP_VERSION=0.2.0 # override the version in the file name
```

| Output in `bin/` | Description |
| --- | --- |
| `wslc-desktop-amd64.exe` | amd64 executable |
| `wslc-desktop-arm64.exe` | arm64 executable |
| `wslc-desktop-0.1.0-amd64.zip` | portable ZIP, contains `wslc-desktop.exe` (amd64) |
| `wslc-desktop-0.1.0-arm64.zip` | portable ZIP, contains `wslc-desktop.exe` (arm64) |

The version in the ZIP name comes from `APP_VERSION` in `Taskfile.yml`; keep it in sync with
`info.version` in `build/config.yml`.

## Project structure

```text
main.go                Entry point: wires the executor and services into Wails
internal/executor/     Command execution layer around wslc / wsl
internal/models/       Data models matching the wslc JSON output
internal/services/     Business logic, one file per resource
build/                 Wails build configuration (Windows only) and icon assets
ui/src/bindings/       Generated TypeScript bindings (git-ignored)
ui/src/api/            Thin typed wrappers around the bindings
ui/src/stores/         Pinia stores, one per resource
ui/src/views/          System, Containers, Images, Volumes, Networks, Terminal
docs/                  Additional documentation
```

## FAQ

**Does it require Docker Desktop?**
No. The app drives the `wslc` CLI that ships with WSL 2; Docker Desktop is not needed.

**What are the minimum requirements?**
Windows 10 build 1809 or newer (Windows 11 recommended), WSL 2, the `wslc` CLI on `PATH`, Go 1.25+,
Node.js 20+ and the Wails v3 CLI for development.

**Does it support Windows on ARM?**
Yes. Build with `wails3 task build ARCH=arm64`, or use `build:all` / `package:zip:all` to produce
both amd64 and ARM64 artifacts at once.

**Can I run a `wslc` command that has no dedicated page?**
Yes. The Terminal view runs any of the 34 `wslc` subcommands and shows stdout, stderr and the exit
code.

**Does the app bundle any runtime?**
No. The generated executable depends only on the WebView2 runtime already present on Windows 10 and
11. A ZIP build can be copied to another machine and launched directly.

**Can I change the interface language?**
Yes. Use the language button in the header to switch between English and Chinese at runtime.

## Documentation

- [简体中文文档](docs/README.zh-CN.md)
- [AGENT.md](AGENT.md) — architecture notes and conventions, useful when working on this codebase
- [Wails v3 documentation](https://v3.wails.io/)
- [WSL documentation](https://learn.microsoft.com/windows/wsl/)

## Contributing

- Found a bug? [Open an issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=bug)
- Need a feature? [Open an issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=enhancement)
- Want to contribute code? [Fork the repo](https://github.com/JFeng2048/Wslc-Desktop/fork),
  create a branch, and [open a pull request](https://github.com/JFeng2048/Wslc-Desktop/pulls/new)
- Questions? [Start a discussion](https://github.com/JFeng2048/Wslc-Desktop/discussions)

## License

No open-source license has been declared yet. Until one is added, all rights are reserved by the
author. If you plan to redistribute the code or build on top of it, please
[open an issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new) first.

## Author

**JFeng2048** — [GitHub](https://github.com/JFeng2048) · [JFeng2048@outlook.com](mailto:JFeng2048@outlook.com)
