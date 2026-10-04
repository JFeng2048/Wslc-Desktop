package services

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"wslc-desktop/internal/executor"
	"wslc-desktop/internal/models"
)

// SystemService 提供系统状态检测与系统信息能力
type SystemService struct {
	exec *executor.Executor
}

// NewSystemService 通过构造函数注入依赖
func NewSystemService(exec *executor.Executor) *SystemService {
	return &SystemService{exec: exec}
}

// GetStatus 检测 WSL 是否可用、wslc 守护进程是否运行
func (s *SystemService) GetStatus(ctx context.Context) (*models.SystemStatus, error) {
	status := &models.SystemStatus{}

	// 检测 WSL 是否可用
	if res, err := s.exec.RunBinary(ctx, "wsl", "--version"); err == nil {
		status.WSLAvailable = true
		status.WSLVersion = parseWSLVersion(res.Stdout)
	} else {
		status.Error = err.Error()
	}

	// 检测 wslc 守护进程
	if res, err := s.exec.Run(ctx, "version"); err == nil {
		status.DaemonRunning = true
		status.DaemonVersion = parseDaemonVersion(res.Stdout)
	} else if status.Error == "" {
		status.Error = res.Stderr
	}

	return status, nil
}

// Info 返回系统信息（wslc info --format json）
func (s *SystemService) Info(ctx context.Context) (*models.SystemInfo, error) {
	res, err := s.exec.Run(ctx, "info", "--format", "json")
	if err != nil {
		return nil, fmt.Errorf("system info: %w", err)
	}
	if res.ExitCode != 0 {
		return nil, fmt.Errorf("system info: %s", res.Stderr)
	}
	var info models.SystemInfo
	if err := json.Unmarshal([]byte(res.Stdout), &info); err != nil {
		return nil, fmt.Errorf("system info parse: %w", err)
	}
	return &info, nil
}

// Version 返回 wslc 版本字符串
func (s *SystemService) Version(ctx context.Context) (string, error) {
	res, err := s.exec.Run(ctx, "version")
	if err != nil {
		return "", fmt.Errorf("version: %w", err)
	}
	return strings.TrimSpace(res.Stdout), nil
}

// versionNumRe 匹配形如 3.0.1.0 / 10.0.26100.9457 的版本号（至少两段数字）
var versionNumRe = regexp.MustCompile(`\d+(?:\.\d+)+`)

// parseWSLVersion 从 `wsl --version` 输出中提取 WSL 自身的版本号。
// 该命令各字段的标签会随系统语言变化（英文 "WSL version:"、中文 "WSL 版本:"），
// 且新版输出含 WSL/WSLg/MSRDC/Direct3D/DXCore/Windows/Kernel 等多行，
// 因此不依赖标签文本，直接取第一行（WSL 本身）中的首个版本号。
func parseWSLVersion(out string) string {
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if m := versionNumRe.FindString(line); m != "" {
			return m
		}
		break
	}
	// 解析失败时返回空字符串，绝不回退输出整段原始内容
	return ""
}

func parseDaemonVersion(out string) string {
	if m := regexp.MustCompile(`(?i)(?:wslc|version)[:\s]+([0-9.]+)`).FindStringSubmatch(out); m != nil {
		return strings.TrimSpace(m[1])
	}
	return strings.TrimSpace(out)
}
