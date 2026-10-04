<div align="center">

# Wslc Desktop

**A graphical desktop manager for WSL containers on Windows.**

[![Platform](https://img.shields.io/badge/platform-Windows-0078D4?logo=windows&logoColor=white)](https://learn.microsoft.com/windows/wsl/)
[![Wails](https://img.shields.io/badge/Wails-v3-06B6D4?logo=go&logoColor=white)](https://v3.wails.io/)
[![Vue](https://img.shields.io/badge/Vue-3.5-42b883?logo=vuedotjs&logoColor=white)](https://vuejs.org/)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev/)
![License](https://img.shields.io/badge/license-unlicensed-lightgrey)

**English** | [简体中文](docs/README.zh-CN.md)

</div>

---

## About

**Wslc Desktop** wraps Microsoft's [`wslc`](https://learn.microsoft.com/windows/wsl/) command-line
interface in a fast, native Windows desktop app. Instead of memorising dozens of subcommands and
flags, you get a clean, searchable UI for every day container operations.

All container work is delegated to the official `wslc` CLI rather than talking to the Docker or WSL
daemon directly, so behaviour always matches what you would get in a terminal.

## Features

| Area | Capabilities |
| --- | --- |
| **System** | WSL availability and version, `wslc` daemon status, version and session manager info |
| **Containers** | List, filter by state, search, start, stop, restart, kill, force-remove, prune, inspect, logs, live stats |
| **Images** | List, search, pull, push, retag, inspect, remove, prune |
| **Volumes** | List, search, create, inspect, remove, prune |
| **Networks** | List, search, create, inspect, remove, prune |
| **Terminal** | Run any of the 34 `wslc` subcommands with stdout / stderr / exit code |

Additional details:

- **Detail drawer** – raw `inspect`, `logs` and `stats` output in a reusable, copy-friendly drawer.
- **Command terminal** – a built-in console for the long tail of `wslc` subcommands that have no
  dedicated page.
- **Light theme only** – a single, carefully tuned light palette for maximum legibility.
- **English & Chinese** – full interface localisation, switchable at runtime.
- **Search and filter** – keyword search plus state filtering on every list view.
- **Responsive tables** – horizontally scrollable data tables with pinned action columns.

## Screenshots

| Volumes | Images | Networks |
| --- | --- | --- |
| ![Volumes view](docs/images/docs-volumes.png) | ![Images view](docs/images/docs-image.png) | ![Networks view](docs/images/docs-network.png) |

## Tech Stack

| Layer | Technology |
| --- | --- |
| Desktop shell | [Wails v3](https://v3.wails.io/) (WebView2 on Windows) |
| Backend | Go 1.25 |
| Frontend | Vue 3 · TypeScript · Vite |
| State | [Pinia](https://pinia.vuejs.org/) |
| Routing | [vue-router](https://router.vuejs.org/) |
| UI kit | [Naive UI](https://www.naiveui.com/) + [Lucide](https://lucide.dev/) icons |
| i18n | [vue-i18n](https://vue-i18n.intlify.dev/) |
| Task runner | [go-task](https://taskfile.dev/) |
| Packaging | NSIS installer and MSIX package |

## Requirements

- Windows 10 (build 1809+) or Windows 11
- WSL 2 with the [`wslc`](https://learn.microsoft.com/windows/wsl/) container CLI available on `PATH`
- [Go](https://go.dev/dl/) 1.25 or newer
- [Node.js](https://nodejs.org/) 20 or newer (ships with npm)

Optional, only for the tooling and packaging commands:

- [Wails v3 CLI](https://v3.wails.io/getting-started/installation/):
  `go install github.com/wailsapp/wails/v3/cmd/wails3@latest`
- [go-task](https://taskfile.dev/docs/installation):
  `go install github.com/go-task/task/v3/cmd/task@latest`
- [NSIS](http://nsis.sourceforge.net/Main_Page) – only needed for NSIS installers
- MSIX packaging additionally needs the MSIX tooling: `task install:msix:tools`

## Getting Started

```bash
# 1. Clone the repository
git clone https://github.com/JFeng2048/Wslc-Desktop.git
cd Wslc-Desktop

# 2. Install frontend dependencies
cd ui && npm install && cd ..

# 3. Run in development mode (hot reload for both Go and frontend)
wails3 dev
```

The Vite dev server and the desktop window start automatically; Go and TypeScript changes are
picked up on save.

> The generated TypeScript bindings in `ui/src/bindings/` are produced by
> `wails3 generate bindings`, which runs automatically as part of `wails3 dev` and `wails3 build`.
> They are git-ignored and must never be edited by hand.

## Common Tasks

All commands below are [go-task](https://taskfile.dev/) tasks. The plain `wails3` equivalents are
given where applicable.

| Command | Wails equivalent | Description |
| --- | --- | --- |
| `task dev` | `wails3 dev` | Start the app with hot reloading |
| `task build` | `wails3 build` | Build `bin/wslc-desktop.exe` |
| `task run` | — | Run the previously built executable |
| `task package` | — | Build an NSIS installer (default) |
| `task install:msix:tools` | — | Install the MSIX packaging toolchain |

Useful variations:

```bash
task package INSTALL_SCOPE=user            # per-user NSIS installer
task package FORMAT=msix                   # MSIX package (publisher CN=JFeng2048)
task build DEV=true                        # development build
```

## Building and Packaging

```bash
# Production build -> bin/wslc-desktop.exe
task build

# NSIS installer
task package

# MSIX package
task package FORMAT=msix
```

The application icon is generated from `build/logo.png`:

```bash
# Regenerate build/windows/icon.ico from the source logo
wails3 generate icons -input build/logo.png -windowsfilename build/windows/icon.ico
```

Signing is optional. Configure a certificate through the Wails CLI (`wails3 setup signing`) or set
`SIGN_CERTIFICATE` / `SIGN_THUMBPRINT` when invoking the `sign` and `sign:installer` tasks.

## Project Structure

```text
main.go                    Entry point: wires the executor and services into Wails
internal/
  executor/                Command execution layer around wslc / wsl
  models/                  Data models matching the wslc JSON output
  services/                Business logic, one file per resource
build/                     Wails build configuration (Windows only) and icon assets
ui/
  src/
    bindings/              Generated TypeScript bindings (git-ignored)
    api/                   Thin typed wrappers around the bindings
    stores/                Pinia stores, one per resource
    views/                 System, Containers, Images, Volumes, Networks, Terminal
    layouts/               Sidebar navigation and header
    i18n/                  English and Chinese message catalogues
```

## Architecture Notes

- Every list is fetched with `--format json` and parsed by a shared, case-insensitive decoder, so
  PascalCase keys from `wslc` and JSON Lines output are both handled.
- `wslc` does not support Go templates for `--format`, which is why JSON is used throughout.
- Arbitrary commands executed from the Terminal view are bounded by a 60 second timeout, and
  long-running follow-style flags are intentionally not exposed, so the UI can never block.
- Frontend types are re-exported from the generated bindings, keeping the Go and TypeScript models
  in sync automatically.

## Contributing

Contributions are welcome. Please open an issue first to discuss larger changes so we can align on
direction before you invest time in an implementation.

- **Found a bug?** [Open an issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=bug)
- **Request a feature?** [Open an issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new?labels=enhancement)
- **Want to contribute code?**
  1. [Fork the repository](https://github.com/JFeng2048/Wslc-Desktop/fork)
  2. Create a topic branch: `git checkout -b feat/my-change`
  3. Commit with clear, atomic commits — English or Chinese, or both
  4. Push the branch to your fork
  5. [Open a pull request](https://github.com/JFeng2048/Wslc-Desktop/pulls/new)
- **Questions or ideas?** [Start a discussion](https://github.com/JFeng2048/Wslc-Desktop/discussions)

When you touch the Go backend, remember to regenerate the bindings before committing:

```bash
wails3 generate bindings -clean=true -ts -i -names -d ui/src/bindings
```

## Issues and Support

- Bug reports and feature requests: [Issues](https://github.com/JFeng2048/Wslc-Desktop/issues)
- Questions: [Discussions](https://github.com/JFeng2048/Wslc-Desktop/discussions)
- Upstream `wslc` problems: [WSL documentation](https://learn.microsoft.com/windows/wsl/)
- Framework questions: [Wails v3 documentation](https://v3.wails.io/) and
  [Wails community](https://github.com/wailsapp/wails/discussions)

## License

No open-source license has been declared for this project yet. Until one is added, all rights are
reserved by the author. If you intend to redistribute the code or build on top of it, please
[open an issue](https://github.com/JFeng2048/Wslc-Desktop/issues/new) first.

## Acknowledgements

- [Wails](https://wails.io/) — the application framework
- [Microsoft `wslc`](https://learn.microsoft.com/windows/wsl/) — the container CLI being managed
- [Naive UI](https://www.naiveui.com/), [Lucide](https://lucide.dev/), [Pinia](https://pinia.vuejs.org/),
  [vue-i18n](https://vue-i18n.intlify.dev/) — UI, icon, state and i18n libraries
- [go-task](https://taskfile.dev/) — task runner used by the build pipeline

## Author

**JFeng2048** — [GitHub](https://github.com/JFeng2048) · [JFeng2048@outlook.com](mailto:JFeng2048@outlook.com)
