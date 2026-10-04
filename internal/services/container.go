package services

import (
	"context"
	"fmt"
	"strings"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/models"
)

// ContainerService 容器管理：列表、启动、停止、重启、删除、清理及详情
type ContainerService struct {
	exec *executor.Executor
}

// NewContainerService 通过构造函数注入依赖
func NewContainerService(exec *executor.Executor) *ContainerService {
	return &ContainerService{exec: exec}
}

// List 列出所有容器（含已停止），使用 --format json 稳定解析
func (s *ContainerService) List(ctx context.Context) ([]models.Container, error) {
	res, err := s.exec.Run(ctx, "list", "-a", "--format", "json")
	if err != nil {
		return nil, fmt.Errorf("list containers: %w", err)
	}
	if res.ExitCode != 0 {
		return nil, fmt.Errorf("list containers: %s", res.Stderr)
	}
	return decodeJSONLines[models.Container](res.Stdout)
}

// Start 启动容器
func (s *ContainerService) Start(ctx context.Context, id string) error {
	return s.run(ctx, "start", id, "start container")
}

// Stop 停止容器
func (s *ContainerService) Stop(ctx context.Context, id string) error {
	return s.run(ctx, "stop", id, "stop container")
}

// Restart 重启容器
func (s *ContainerService) Restart(ctx context.Context, id string) error {
	return s.run(ctx, "restart", id, "restart container")
}

// Kill 强制终止容器
func (s *ContainerService) Kill(ctx context.Context, id string) error {
	return s.run(ctx, "kill", id, "kill container")
}

// Remove 删除容器；force 为 true 时追加 -f
func (s *ContainerService) Remove(ctx context.Context, id string, force bool) error {
	args := []string{"rm"}
	if force {
		args = append(args, "-f")
	}
	args = append(args, id)
	return s.runArgs(ctx, args, fmt.Sprintf("remove container %s", id))
}

// Prune 清理所有已停止容器，返回命令输出
func (s *ContainerService) Prune(ctx context.Context) (string, error) {
	res, err := s.exec.Run(ctx, "container", "prune", "-f")
	if err != nil {
		return "", fmt.Errorf("prune containers: %w", err)
	}
	return strings.TrimSpace(res.Stdout + res.Stderr), nil
}

// Inspect 返回容器详情的原始 JSON
func (s *ContainerService) Inspect(ctx context.Context, id string) (string, error) {
	return s.output(ctx, "inspect", id)
}

// Logs 返回容器日志；tail>0 时限制行数，timestamps 时附加时间戳
func (s *ContainerService) Logs(ctx context.Context, id string, tail int, timestamps bool) (string, error) {
	args := []string{"logs"}
	if tail > 0 {
		args = append(args, "--tail", fmt.Sprintf("%d", tail))
	}
	if timestamps {
		args = append(args, "--timestamps")
	}
	args = append(args, id)
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return "", fmt.Errorf("logs: %w", err)
	}
	return res.Stdout + res.Stderr, nil
}

// Stats 返回容器资源使用快照（JSON）
func (s *ContainerService) Stats(ctx context.Context) (string, error) {
	res, err := s.exec.Run(ctx, "stats", "--format", "json")
	if err != nil {
		return "", fmt.Errorf("stats: %w", err)
	}
	return res.Stdout, nil
}

func (s *ContainerService) run(ctx context.Context, action, id, op string) error {
	return s.runArgs(ctx, []string{action, id}, op)
}

func (s *ContainerService) runArgs(ctx context.Context, args []string, op string) error {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("%s failed: %s", op, res.Stderr)
	}
	return nil
}

func (s *ContainerService) output(ctx context.Context, args ...string) (string, error) {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return "", fmt.Errorf("%v: %w", args, err)
	}
	if res.ExitCode != 0 {
		return "", fmt.Errorf("%v failed: %s", args, res.Stderr)
	}
	return res.Stdout, nil
}
