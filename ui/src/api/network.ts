import { NetworkService } from "@/bindings/wslc-desktop/internal/services";
import type { Network } from "@/types";

export function listNetworks(): Promise<Network[]> {
  return NetworkService.List() as unknown as Promise<Network[]>;
}
export function removeNetwork(id: string): Promise<void> {
  return NetworkService.Remove(id) as unknown as Promise<void>;
}
export function pruneNetworks(): Promise<string> {
  return NetworkService.Prune() as unknown as Promise<string>;
}
export function inspectNetwork(id: string): Promise<string> {
  return NetworkService.Inspect(id) as unknown as Promise<string>;
}
export function createNetwork(name: string): Promise<void> {
  return NetworkService.Create(name) as unknown as Promise<void>;
}
