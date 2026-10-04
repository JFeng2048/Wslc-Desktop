package executor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
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
