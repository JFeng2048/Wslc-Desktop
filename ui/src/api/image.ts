import { ImageService } from "@/bindings/wslc-desktop/internal/services";
import type { Image } from "@/types";

export function listImages(): Promise<Image[]> {
  return ImageService.List() as unknown as Promise<Image[]>;
}
export function removeImage(id: string, force: boolean): Promise<void> {
  return ImageService.Remove(id, force) as unknown as Promise<void>;
}
export function inspectImage(id: string): Promise<string> {
  return ImageService.Inspect(id) as unknown as Promise<string>;
}
export function pullImage(ref: string): Promise<string> {
  return ImageService.Pull(ref) as unknown as Promise<string>;
}
export function pushImage(ref: string): Promise<string> {
  return ImageService.Push(ref) as unknown as Promise<string>;
}
export function tagImage(source: string, target: string): Promise<void> {
  return ImageService.Tag(source, target) as unknown as Promise<void>;
}
export function pruneImages(): Promise<string> {
  return ImageService.Prune() as unknown as Promise<string>;
}
