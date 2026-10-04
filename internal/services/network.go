package services

import (
	"context"
	"fmt"
	"strings"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/models"
)

// NetworkService 网络管理：列表、删除、清理、详情、创建
type NetworkService struct {
	exec *executor.Executor
}

// NewNetworkService 通过构造函数注入依赖
func NewNetworkService(exec *executor.Executor) *NetworkService {
	return &NetworkService{exec: exec}
}

// List 列出网络，使用 --format json 稳定解析
func (s *NetworkService) List(ctx context.Context) ([]models.Network, error) {
	res, err := s.exec.Run(ctx, "network", "ls", "--format", "json")
	if err != nil {
		return nil, fmt.Errorf("list networks: %w", err)
	}
	if res.ExitCode != 0 {
		return nil, fmt.Errorf("list networks: %s", res.Stderr)
	}
	return decodeJSONLines[models.Network](res.Stdout)
}

// Remove 删除网络
func (s *NetworkService) Remove(ctx context.Context, id string) error {
	return s.runArgs(ctx, []string{"network", "rm", id}, fmt.Sprintf("remove network %s", id))
}

// Prune 清理未使用网络，返回命令输出
func (s *NetworkService) Prune(ctx context.Context) (string, error) {
	res, err := s.exec.Run(ctx, "network", "prune", "-f")
	if err != nil {
		return "", fmt.Errorf("prune networks: %w", err)
	}
	return strings.TrimSpace(res.Stdout + res.Stderr), nil
}

// Inspect 返回网络详情的原始 JSON
func (s *NetworkService) Inspect(ctx context.Context, id string) (string, error) {
	return s.output(ctx, "inspect", id)
}

// Create 创建网络
func (s *NetworkService) Create(ctx context.Context, name string) error {
	return s.runArgs(ctx, []string{"network", "create", name}, fmt.Sprintf("create network %s", name))
}

func (s *NetworkService) runArgs(ctx context.Context, args []string, op string) error {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("%s failed: %s", op, res.Stderr)
	}
	return nil
}

func (s *NetworkService) output(ctx context.Context, args ...string) (string, error) {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return "", fmt.Errorf("%v: %w", args, err)
	}
	if res.ExitCode != 0 {
		return "", fmt.Errorf("%v failed: %s", args, res.Stderr)
	}
	return res.Stdout, nil
}
