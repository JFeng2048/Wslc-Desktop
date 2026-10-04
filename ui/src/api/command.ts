import { CommandService } from "@/bindings/wslc-desktop/internal/services";
import type { CommandResult } from "@/types";

// 运行任意 wslc 子命令，覆盖全部 34 个子命令
export function runCommand(args: string[]): Promise<CommandResult> {
  return CommandService.Run(args) as unknown as Promise<CommandResult>;
}
