import { defineStore } from "pinia";
import { ref } from "vue";
import * as api from "@/api/network";
import type { Network } from "@/types";

export const useNetworkStore = defineStore("network", () => {
  const list = ref<Network[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  async function fetch() {
    loading.value = true;
    error.value = null;
    try {
      list.value = await api.listNetworks();
    } catch (e) {
      error.value = (e as Error)?.message ?? String(e);
    } finally {
      loading.value = false;
    }
  }

  const remove = (id: string) => api.removeNetwork(id);
  const prune = () => api.pruneNetworks();
  const inspect = (id: string) => api.inspectNetwork(id);
  const create = (name: string) => api.createNetwork(name);

  return { list, loading, error, fetch, remove, prune, inspect, create };
});
