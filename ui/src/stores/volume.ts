import { defineStore } from "pinia";
import { ref } from "vue";
import * as api from "@/api/volume";
import type { Volume } from "@/types";

export const useVolumeStore = defineStore("volume", () => {
  const list = ref<Volume[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  async function fetch() {
    loading.value = true;
    error.value = null;
    try {
      list.value = await api.listVolumes();
    } catch (e) {
      error.value = (e as Error)?.message ?? String(e);
    } finally {
      loading.value = false;
    }
  }

  const remove = (name: string) => api.removeVolume(name);
  const prune = () => api.pruneVolumes();
  const inspect = (name: string) => api.inspectVolume(name);
  const create = (name: string) => api.createVolume(name);

  return { list, loading, error, fetch, remove, prune, inspect, create };
});
