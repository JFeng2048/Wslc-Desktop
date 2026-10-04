import { defineStore } from "pinia";
import { ref } from "vue";
import * as api from "@/api/image";
import type { Image } from "@/types";

export const useImageStore = defineStore("image", () => {
  const list = ref<Image[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);

  async function fetch() {
    loading.value = true;
    error.value = null;
    try {
      list.value = await api.listImages();
    } catch (e) {
      error.value = (e as Error)?.message ?? String(e);
    } finally {
      loading.value = false;
    }
  }

  const remove = (id: string, force: boolean) => api.removeImage(id, force);
  const inspect = (id: string) => api.inspectImage(id);
  const pull = (ref: string) => api.pullImage(ref);
  const push = (ref: string) => api.pushImage(ref);
  const tag = (source: string, target: string) => api.tagImage(source, target);
  const prune = () => api.pruneImages();

  return { list, loading, error, fetch, remove, inspect, pull, push, tag, prune };
});
