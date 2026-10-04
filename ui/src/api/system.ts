import { SystemService } from "@/bindings/wslc-desktop/internal/services";
import type { SystemStatus, SystemInfo } from "@/types";

// 检测 WSL 是否可用、wslc 守护进程是否运行
export function getSystemStatus(): Promise<SystemStatus> {
  return SystemService.GetStatus() as unknown as Promise<SystemStatus>;
}
// 系统信息（wslc info --format json）
export function getSystemInfo(): Promise<SystemInfo> {
  return SystemService.Info() as unknown as Promise<SystemInfo>;
}
// wslc 版本
export function getVersion(): Promise<string> {
  return SystemService.Version() as unknown as Promise<string>;
}
