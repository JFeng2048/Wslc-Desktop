<script setup lang="ts">
import { onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { storeToRefs } from "pinia";
import { NGrid, NGi, NCard, NTag, NButton, NAlert, NEmpty, NSpin } from "naive-ui";
import { RefreshCw } from "lucide-vue-next";
import { useSystemStore } from "@/stores/system";

const { t } = useI18n();
const store = useSystemStore();
const { status, loading, error } = storeToRefs(store);

onMounted(store.fetch);

function yn(v: boolean) {
  return v ? t("common.yes") : t("common.no");
}
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h2 class="page-title">{{ t("system.title") }}</h2>
      <n-button :loading="loading" secondary strong @click="store.fetch()">
        <template #icon><RefreshCw :size="16" /></template>
        {{ t("common.refresh") }}
      </n-button>
    </div>

    <p class="page-tip">{{ t("system.tip") }}</p>

    <n-alert
      v-if="error"
      type="error"
      :title="t('system.detectError')"
      class="alert-block"
    >
      {{ error }}
    </n-alert>

    <n-spin v-if="loading && !status" :show="loading">
      <div style="min-height: 200px" />
    </n-spin>

    <n-grid
      v-else-if="status"
      :cols="2"
      :x-gap="16"
      :y-gap="16"
      responsive="screen"
      item-responsive
    >
      <n-gi span="24 m:12">
        <n-card :title="t('system.wslAvailable')">
          <n-tag :type="status.wslAvailable ? 'success' : 'error'">
            {{ yn(status.wslAvailable) }}
          </n-tag>
          <div class="card-meta">{{ status.wslVersion || t("system.notAvailable") }}</div>
        </n-card>
      </n-gi>
      <n-gi span="24 m:12">
        <n-card :title="t('system.daemon')">
          <n-tag :type="status.daemonRunning ? 'success' : 'warning'">
            {{ yn(status.daemonRunning) }}
          </n-tag>
          <div class="card-meta">
            {{ status.daemonVersion || t("system.notRunning") }}
          </div>
        </n-card>
      </n-gi>
    </n-grid>

    <n-empty v-else :description="t('common.empty')" />
  </div>
</template>
