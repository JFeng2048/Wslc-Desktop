package executor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// Result 保存一次命令执行的结果
type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

// Executor 统一封装 wslc 及其依赖命令（如 wsl）的调用。
// Service 层不直接使用 exec.Command，全部走这里。
type Executor struct {
	binary string
}

// New 创建 Executor；binary 为空时默认使用 "wslc"
func New(binary string) *Executor {
	if binary == "" {
		binary = "wslc"
	}
	return &Executor{binary: binary}
}

// Run 执行一条 wslc 命令
func (e *Executor) Run(ctx context.Context, args ...string) (*Result, error) {
	return e.RunBinary(ctx, e.binary, args...)
}

// RunBinary 执行任意外部命令（如 wsl），用于系统能力探测等场景
func (e *Executor) RunBinary(ctx context.Context, binary string, args ...string) (*Result, error) {
	cmd := exec.CommandContext(ctx, binary, args...)

	// 桌面程序没有控制台，若不隐藏窗口，每次执行命令都会闪出一个 cmd 窗口
	hideCmdWindow(cmd)

	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	res := &Result{
		Stdout: strings.TrimSpace(out.String()),
		Stderr: strings.TrimSpace(errOut.String()),
	}

	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// 命令执行了但返回非零退出码：结果仍可用，交由上层业务判断
			res.ExitCode = exitErr.ExitCode()
			return res, nil
		}
		// 命令本身无法启动（例如二进制不存在）
		return res, fmt.Errorf("exec %q failed: %w", binary, err)
	}
	res.ExitCode = 0
	return res, nil
}

// createNoWindow 为 Win32 CREATE_NO_WINDOW 标志值。
// Go 的 syscall 包未导出该常量，这里按 Win32 定义直接声明，避免引入额外依赖。
const createNoWindow = 0x08000000

// hideCmdWindow 隐藏子进程的控制台窗口。
//
// 本项目以 windowsgui 子系统构建，主程序本身不附带控制台；若子进程沿用默认
// 设置，每执行一条命令（例如 wslc list）屏幕上都会闪现一个 cmd 窗口。
// CREATE_NO_WINDOW 让子进程在无控制台的情况下静默启动，HideWindow 兜底。
func hideCmdWindow(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.HideWindow = true
	cmd.SysProcAttr.CreationFlags |= createNoWindow
}
