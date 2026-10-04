import { defineStore } from "pinia";
import { ref } from "vue";
import { getSystemStatus, getSystemInfo, getVersion } from "@/api/system";
import { toMessage } from "@/utils/error";
import type { SystemStatus, SystemInfo } from "@/types";

export const useSystemStore = defineStore("system", () => {
  const status = ref<SystemStatus | null>(null);
  const info = ref<SystemInfo | null>(null);
  const version = ref<string>("");
  const loading = ref(false);
  const error = ref<string | null>(null);

  async function fetch() {
    loading.value = true;
    error.value = null;
    try {
      status.value = await getSystemStatus();
      info.value = await getSystemInfo();
      version.value = await getVersion();
    } catch (e) {
      error.value = toMessage(e);
    } finally {
      loading.value = false;
    }
  }

  return { status, info, version, loading, error, fetch };
});
