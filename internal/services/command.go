package services

import (
	"context"
	"fmt"
	"time"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/models"
)

// CommandService 通用命令执行器：可在前端命令终端中运行任意 wslc 子命令，
// 从而覆盖全部 34 个子命令（attach/build/create/exec/events/export/import/
// load/login/logout/run/save/tag 等未在结构化页面中单独提供的命令）。
type CommandService struct {
	exec *executor.Executor
}

// NewCommandService 通过构造函数注入依赖
func NewCommandService(exec *executor.Executor) *CommandService {
	return &CommandService{exec: exec}
}

// Run 执行一个 wslc 子命令（args[0] 为子命令，如 ["logs", "myc", "--tail", "50"]）。
// 为避免 follow(-f) 等命令无限阻塞，这里统一加上 60s 超时。返回合并后的输出与退出码。
func (s *CommandService) Run(ctx context.Context, args []string) (*models.CommandResult, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("no command provided")
	}
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("run %v: %w", args, err)
	}
	return &models.CommandResult{
		Stdout:   res.Stdout,
		Stderr:   res.Stderr,
		ExitCode: res.ExitCode,
	}, nil
}
