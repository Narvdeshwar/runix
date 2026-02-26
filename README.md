# MyTerm Studio

A high-end, immersive Terminal User Interface (TUI) for managing containerized workspaces and isolated runtime environments directly from your CLI. Designed with a modern, cyberpunk-inspired control panel aesthetic.

## Features

- **Cyberpunk Aesthetic**: An intense, visually striking console wrapped in neon and dark mode styling.
- **Isolated Workspaces**: Effortlessly launch sandboxed Linux containers directly connected to your host.
- **Interactive Control Panel**: Navigate environments and dive into system configurations via intuitive keyboard bindings.
- **Built in Go**: Fast, independent executable powered by the robust `bubbletea` and `lipgloss` libraries.

---

## 🚀 Getting Started

### Prerequisites
- [Go 1.25+](https://go.dev/doc/install) installed on your system.
- PowerShell or Command Prompt (Windows), or Bash (Linux/macOS).

### 1. Build the Executable
Compile the source code into a standalone binary:

```bash
go build -o myterm.exe .
```

### 2. Launching the Dashboard (TUI)
To open the modern control panel interface (opens elegantly in a new native window on Windows):

```bash
./myterm.exe dashboard
```

### 3. Direct Shell Access
If you want to instantly drop into the isolated workspace without seeing the dashboard navigation first:

```bash
./myterm.exe shell
```

---

## ⌨️ TUI Keyboard Controls

When inside the `dashboard` interface, use the following keys:
- `↑` / `k` : Move selection pointer up
- `↓` / `j` : Move selection pointer down
- `Enter` / `Space` : Initiate the boot sequence for the selected workspace
- `c` : Toggle System Configuration overly
- `q` / `Ctrl+C` : Terminate and exit the interface completely

---

## 🛠 Project Structure

- `/cmd` - Cobra CLI commands parsing (root, dashboard, shell).
- `/internal/tui` - Contains all visual logic (`app.go`) and custom theme styling (`styles.go`).
- `/internal/runtime` - Workspace daemon, container runtime logic.
- `/internal/setup` - Application bootstrapping and data dir prep.
