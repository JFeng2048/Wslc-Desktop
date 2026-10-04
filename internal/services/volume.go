package services

import (
	"context"
	"fmt"
	"strings"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/models"
)

// VolumeService 卷管理：列表、删除、清理、详情、创建
type VolumeService struct {
	exec *executor.Executor
}

// NewVolumeService 通过构造函数注入依赖
func NewVolumeService(exec *executor.Executor) *VolumeService {
	return &VolumeService{exec: exec}
}

// List 列出卷，使用 --format json 稳定解析
func (s *VolumeService) List(ctx context.Context) ([]models.Volume, error) {
	res, err := s.exec.Run(ctx, "volume", "ls", "--format", "json")
	if err != nil {
		return nil, fmt.Errorf("list volumes: %w", err)
	}
	if res.ExitCode != 0 {
		return nil, fmt.Errorf("list volumes: %s", res.Stderr)
	}
	return decodeJSONLines[models.Volume](res.Stdout)
}

// Remove 删除卷
func (s *VolumeService) Remove(ctx context.Context, name string) error {
	return s.runArgs(ctx, []string{"volume", "rm", name}, fmt.Sprintf("remove volume %s", name))
}

// Prune 清理未使用卷，返回命令输出
func (s *VolumeService) Prune(ctx context.Context) (string, error) {
	res, err := s.exec.Run(ctx, "volume", "prune", "-f")
	if err != nil {
		return "", fmt.Errorf("prune volumes: %w", err)
	}
	return strings.TrimSpace(res.Stdout + res.Stderr), nil
}

// Inspect 返回卷详情的原始 JSON
func (s *VolumeService) Inspect(ctx context.Context, name string) (string, error) {
	return s.output(ctx, "inspect", name)
}

// Create 创建卷
func (s *VolumeService) Create(ctx context.Context, name string) error {
	return s.runArgs(ctx, []string{"volume", "create", name}, fmt.Sprintf("create volume %s", name))
}

func (s *VolumeService) runArgs(ctx context.Context, args []string, op string) error {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("%s failed: %s", op, res.Stderr)
	}
	return nil
}

func (s *VolumeService) output(ctx context.Context, args ...string) (string, error) {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return "", fmt.Errorf("%v: %w", args, err)
	}
	if res.ExitCode != 0 {
		return "", fmt.Errorf("%v failed: %s", args, res.Stderr)
	}
	return res.Stdout, nil
}
