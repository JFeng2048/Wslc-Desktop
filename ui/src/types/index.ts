// 与后端 internal/models 字段保持一致（由 wails 绑定生成，统一从此处引用避免漂移）
export type {
  Container,
  Image,
  Volume,
  Network,
  SystemStatus,
  SystemInfo,
  SystemClient,
  SystemServer,
  SessionInfo,
  CommandResult,
} from "@/bindings/wslc-desktop/internal/models";
