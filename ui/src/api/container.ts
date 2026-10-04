import { ContainerService } from "@/bindings/wslc-desktop/internal/services";
import type { Container } from "@/types";

export function listContainers(): Promise<Container[]> {
  return ContainerService.List() as unknown as Promise<Container[]>;
}
export function startContainer(id: string): Promise<void> {
  return ContainerService.Start(id) as unknown as Promise<void>;
}
export function stopContainer(id: string): Promise<void> {
  return ContainerService.Stop(id) as unknown as Promise<void>;
}
export function restartContainer(id: string): Promise<void> {
  return ContainerService.Restart(id) as unknown as Promise<void>;
}
export function killContainer(id: string): Promise<void> {
  return ContainerService.Kill(id) as unknown as Promise<void>;
}
export function removeContainer(id: string, force: boolean): Promise<void> {
  return ContainerService.Remove(id, force) as unknown as Promise<void>;
}
export function pruneContainers(): Promise<string> {
  return ContainerService.Prune() as unknown as Promise<string>;
}
export function inspectContainer(id: string): Promise<string> {
  return ContainerService.Inspect(id) as unknown as Promise<string>;
}
export function containerLogs(id: string, tail: number, timestamps: boolean): Promise<string> {
  return ContainerService.Logs(id, tail, timestamps) as unknown as Promise<string>;
}
export function containerStats(): Promise<string> {
  return ContainerService.Stats() as unknown as Promise<string>;
}
