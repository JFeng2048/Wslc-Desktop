package services

import (
	"context"
	"fmt"
	"strings"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/models"
)

// ImageService 镜像管理：列表、删除、详情、拉取、推送、打标签、清理
type ImageService struct {
	exec *executor.Executor
}

// NewImageService 通过构造函数注入依赖
func NewImageService(exec *executor.Executor) *ImageService {
	return &ImageService{exec: exec}
}

// List 列出本地镜像，使用 --format json 稳定解析
func (s *ImageService) List(ctx context.Context) ([]models.Image, error) {
	res, err := s.exec.Run(ctx, "images", "--format", "json")
	if err != nil {
		return nil, fmt.Errorf("list images: %w", err)
	}
	if res.ExitCode != 0 {
		return nil, fmt.Errorf("list images: %s", res.Stderr)
	}
	return decodeJSONLines[models.Image](res.Stdout)
}

// Remove 删除镜像；force 为 true 时追加 -f
func (s *ImageService) Remove(ctx context.Context, id string, force bool) error {
	args := []string{"rmi"}
	if force {
		args = append(args, "-f")
	}
	args = append(args, id)
	return s.runArgs(ctx, args, fmt.Sprintf("remove image %s", id))
}

// Inspect 返回镜像详情的原始 JSON
func (s *ImageService) Inspect(ctx context.Context, id string) (string, error) {
	return s.output(ctx, "inspect", id)
}

// Pull 拉取镜像，返回命令输出
func (s *ImageService) Pull(ctx context.Context, ref string) (string, error) {
	return s.output(ctx, "pull", ref)
}

// Push 推送镜像，返回命令输出
func (s *ImageService) Push(ctx context.Context, ref string) (string, error) {
	return s.output(ctx, "push", ref)
}

// Tag 为镜像打标签（source 如 repo:tag，target 为目标引用）
func (s *ImageService) Tag(ctx context.Context, source, target string) error {
	return s.runArgs(ctx, []string{"tag", source, target}, "tag image")
}

// Prune 清理未使用的镜像，返回命令输出
func (s *ImageService) Prune(ctx context.Context) (string, error) {
	res, err := s.exec.Run(ctx, "image", "prune", "-f")
	if err != nil {
		return "", fmt.Errorf("prune images: %w", err)
	}
	return strings.TrimSpace(res.Stdout + res.Stderr), nil
}

func (s *ImageService) runArgs(ctx context.Context, args []string, op string) error {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("%s failed: %s", op, res.Stderr)
	}
	return nil
}

func (s *ImageService) output(ctx context.Context, args ...string) (string, error) {
	res, err := s.exec.Run(ctx, args...)
	if err != nil {
		return "", fmt.Errorf("%v: %w", args, err)
	}
	if res.ExitCode != 0 {
		return "", fmt.Errorf("%v failed: %s", args, res.Stderr)
	}
	return res.Stdout, nil
}
