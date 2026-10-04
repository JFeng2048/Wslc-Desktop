import { defineStore } from "pinia";
import { ref } from "vue";
import * as api from "@/api/container";
import type { Container } from "@/types";

export const useContainerStore = defineStore("container", () => {
  const list = ref<Container[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  async function fetch() {
    loading.value = true;
    error.value = null;
    try {
      list.value = await api.listContainers();
    } catch (e) {
      error.value = (e as Error)?.message ?? String(e);
    } finally {
      loading.value = false;
    }
  }

  const start = (id: string) => api.startContainer(id);
  const stop = (id: string) => api.stopContainer(id);
  const restart = (id: string) => api.restartContainer(id);
  const kill = (id: string) => api.killContainer(id);
  const remove = (id: string, force: boolean) => api.removeContainer(id, force);
  const prune = () => api.pruneContainers();
  const inspect = (id: string) => api.inspectContainer(id);
  const logs = (id: string, tail: number, ts: boolean) => api.containerLogs(id, tail, ts);
  const stats = () => api.containerStats();

  return { list, loading, error, fetch, start, stop, restart, kill, remove, prune, inspect, logs, stats };
});
