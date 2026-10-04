import { VolumeService } from "@/bindings/wslc-desktop/internal/services";
import type { Volume } from "@/types";

export function listVolumes(): Promise<Volume[]> {
  return VolumeService.List() as unknown as Promise<Volume[]>;
}
export function removeVolume(name: string): Promise<void> {
  return VolumeService.Remove(name) as unknown as Promise<void>;
}
export function pruneVolumes(): Promise<string> {
  return VolumeService.Prune() as unknown as Promise<string>;
}
export function inspectVolume(name: string): Promise<string> {
  return VolumeService.Inspect(name) as unknown as Promise<string>;
}
export function createVolume(name: string): Promise<void> {
  return VolumeService.Create(name) as unknown as Promise<void>;
}
