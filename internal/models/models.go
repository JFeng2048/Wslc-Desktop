package models

import "time"

// Container 容器信息；json 标签为 camelCase，解析时由 services.decodeJSONLines
// 按大小写不敏感匹配 wslc 输出的 PascalCase 键（如 Names / CreatedAt）。
type Container struct {
	ID           string `json:"id"`
	Names        string `json:"names"`
	Image        string `json:"image"`
	Command      string `json:"command"`
	CreatedAt    string `json:"createdAt"`
	RunningFor   string `json:"runningFor"`
	State        string `json:"state"`
	Status       string `json:"status"`
	Ports        string `json:"ports"`
	Size         string `json:"size"`
	Labels       string `json:"labels"`
	Networks     string `json:"networks"`
	Mounts       string `json:"mounts"`
	LocalVolumes string `json:"localVolumes"`
	HealthStatus string `json:"healthStatus"`
}

// Image 本地镜像信息
type Image struct {
	ID           string `json:"id"`
	Repository   string `json:"repository"`
	Tag          string `json:"tag"`
	Digest       string `json:"digest"`
	Size         string `json:"size"`
	CreatedAt    string `json:"createdAt"`
	CreatedSince string `json:"createdSince"`
	Containers   string `json:"containers"`
}

// Volume 卷信息
type Volume struct {
	Name        string `json:"name"`
	Driver      string `json:"driver"`
	Scope       string `json:"scope"`
	Mountpoint  string `json:"mountpoint"`
	Labels      string `json:"labels"`
	Status      string `json:"status"`
	Size        string `json:"size"`
	Links       string `json:"links"`
	Group       string `json:"group"`
	Availability string `json:"availability"`
}

// Network 网络信息
type Network struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Driver    string `json:"driver"`
	Scope     string `json:"scope"`
	CreatedAt string `json:"createdAt"`
	IPv4      bool   `json:"ipv4"`
	IPv6      bool   `json:"ipv6"`
	Internal  bool   `json:"internal"`
	Labels    string `json:"labels"`
}

// SystemStatus 系统运行状态（WSL / wslc 守护进程探测）
type SystemStatus struct {
	WSLAvailable  bool   `json:"wslAvailable"`
	WSLVersion    string `json:"wslVersion"`
	DaemonRunning bool   `json:"daemonRunning"`
	DaemonVersion string `json:"daemonVersion"`
	Error         string `json:"error,omitempty"`
}

// SystemClient / SystemServer / SessionInfo 对应 `wslc info --format json`
type SystemClient struct {
	Direct3DVersion string `json:"direct3DVersion"`
	DxCoreVersion   string `json:"dxCoreVersion"`
	KernelVersion   string `json:"kernelVersion"`
	SettingsFile    string `json:"settingsFile"`
	Version         string `json:"version"`
	WindowsVersion  string `json:"windowsVersion"`
}

// SystemServer 对应 `wslc info` 的 Server 段
type SystemServer struct {
	SessionManagerVersion string        `json:"sessionManagerVersion"`
	Sessions              []SessionInfo `json:"sessions"`
}

// SessionInfo 会话信息
type SessionInfo struct {
	CreatorPid int    `json:"creatorPid"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
}

// SystemInfo 对应 `wslc info --format json` 的完整结构
type SystemInfo struct {
	Client SystemClient `json:"client"`
	Server SystemServer `json:"server"`
}

// CommandResult 通用命令执行结果，供命令终端展示
type CommandResult struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exitCode"`
}

// 保留 time 依赖（前端绑定中 time.Time 表现为字符串）
var _ = time.Time{}
