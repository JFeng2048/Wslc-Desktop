<div align="center">

# Wslc Desktop

**A graphical desktop manager for WSL containers on Windows.**

[![Platform](https://img.shields.io/badge/platform-Windows-0078D4?logo=windows&logoColor=white)](https://learn.microsoft.com/windows/wsl/)
[![Wails](https://img.shields.io/badge/Wails-v3-06B6D4?logo=go&logoColor=white)](https://v3.wails.io/)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
![License](https://img.shields.io/badge/license-unlicensed-lightgrey)

**English** | [简体中文](docs/README.zh-CN.md)

</div>

---

## About

Wslc Desktop wraps Microsoft's [`wslc`](https://learn.microsoft.com/windows/wsl/) CLI into a native
Windows desktop app, so everyday container work needs no command memorisation.

Built with Wails v3, Go and Vue 3.

## Screenshots

| Volumes | Images | Networks |
| --- | --- | --- |
| ![Volumes view](docs/images/docs-volumes.png) | ![Images view](docs/images/docs-image.png) | ![Networks view](docs/images/docs-network.png) |

## Features

- **System** — WSL availability and version, `wslc` daemon status
- **Containers** — list, search, filter by state, start, stop, restart, kill, remove, logs, stats, prune
- **Images** — list, search, pull, push, retag, remove, prune
- **Volumes** — list, search, create, remove, prune
- **Networks** — list, search, create, remove, prune
- **Terminal** — run any `wslc` subcommand and read its output
- English and Chinese interface, light theme

## Requirements

- Windows 10 (build 1809+) or Windows 11
- WSL 2 with the `wslc` CLI available on `PATH`
- [Go](https://go.dev/dl/) 1.25 or newer
- [Node.js](https://nodejs.org/) 20 or newer
- [Wails v3 CLI](https://v3.wails.io/getting-started/installation/):
  `go install github.com/wailsapp/wails/v3/cmd/wails3@latest`

## Usage

Two commands are all you need:

```bash
# Development mode, with hot reload for the frontend and backend
wails3 dev

# Production build -> bin/wslc-desktop-amd64.exe
wails3 build
```

### Build for a specific architecture

The default target is `amd64`. Pass `ARCH` to cross-compile — the architecture is always part of the
binary name, so multiple builds can coexist in `bin/`.

```bash
wails3 task build ARCH=arm64              # -> bin/wslc-desktop-arm64.exe
wails3 task build:all                     # both amd64 and arm64
```

### Package as ZIP

Each ZIP contains the executable named `wslc-desktop.exe`, while the archive name carries the
version and architecture.

```bash
wails3 task package:zip                   # -> bin/wslc-desktop-0.1.0-amd64.zip
wails3 task package:zip:all               # ZIPs for both architectures
wails3 task package:zip APP_VERSION=0.2.0 # override the version in the file name
```

Output in `bin/`:

| File | Description |
| --- | --- |
| `wslc-desktop-amd64.exe` | amd64 executable |
| `wslc-desktop-arm64.exe` | arm64 executable |
| `wslc-desktop-0.1.0-amd64.zip` | portable ZIP, contains `wslc-desktop.exe` (amd64) |
| `wslc-desktop-0.1.0-arm64.zip` | portable ZIP, contains `wslc-desktop.exe` (arm64) |

> The version in the ZIP name comes from `APP_VERSION` in `Taskfile.yml`; keep it in sync with
> `info.version` in `build/config.yml`.

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
