<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { storeToRefs } from "pinia";
import { NGrid, NGi, NCard, NTag, NButton, NAlert, NEmpty, NSpin, NDescriptions, NDescriptionsItem } from "naive-ui";
import { RefreshCw } from "lucide-vue-next";
import { useSystemStore } from "@/stores/system";

const { t } = useI18n();
const store = useSystemStore();
const { status, info, version, loading, error } = storeToRefs(store);

onMounted(store.fetch);

const yn = (v: boolean) => (v ? t("common.yes") : t("common.no"));

const sessions = computed(() => info.value?.server?.sessions ?? []);
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

    <n-alert v-if="error" type="error" :title="t('system.detectError')" class="alert-block">
      {{ error }}
    </n-alert>

    <n-spin v-if="loading && !status" :show="loading">
      <div style="min-height: 200px" />
    </n-spin>

    <template v-else>
      <n-grid :cols="2" :x-gap="16" :y-gap="16" responsive="screen" item-responsive>
        <n-gi span="24 m:12">
          <n-card :title="t('system.wslAvailable')">
            <n-tag :type="status?.wslAvailable ? 'success' : 'error'">{{ yn(status?.wslAvailable ?? false) }}</n-tag>
            <div class="card-meta">{{ status?.wslVersion || t("system.notAvailable") }}</div>
          </n-card>
        </n-gi>
        <n-gi span="24 m:12">
          <n-card :title="t('system.daemon')">
            <n-tag :type="status?.daemonRunning ? 'success' : 'warning'">{{ yn(status?.daemonRunning ?? false) }}</n-tag>
            <div class="card-meta">{{ status?.daemonVersion || t("system.notRunning") }}</div>
          </n-card>
        </n-gi>
        <n-gi span="24 m:12">
          <n-card :title="t('system.version')">
            <div class="card-meta mono">{{ version || "—" }}</div>
          </n-card>
        </n-gi>
        <n-gi span="24 m:12">
          <n-card :title="t('system.sessionManager')">
            <div class="card-meta mono">{{ info?.server?.sessionManagerVersion || "—" }}</div>
          </n-card>
        </n-gi>
      </n-grid>

      <n-card :title="t('system.clientInfo')" class="mt-16">
        <n-descriptions v-if="info?.client" label-placement="left" bordered :column="1">
          <n-descriptions-item :label="t('system.kernel')">{{ info.client.kernelVersion }}</n-descriptions-item>
          <n-descriptions-item :label="t('system.windows')">{{ info.client.windowsVersion }}</n-descriptions-item>
          <n-descriptions-item :label="t('system.direct3d')">{{ info.client.direct3DVersion }}</n-descriptions-item>
          <n-descriptions-item :label="t('system.dxcore')">{{ info.client.dxCoreVersion }}</n-descriptions-item>
          <n-descriptions-item :label="t('system.settingsFile')">{{ info.client.settingsFile }}</n-descriptions-item>
        </n-descriptions>
        <n-empty v-else :description="t('common.empty')" />
      </n-card>

      <n-card :title="t('system.sessions')" class="mt-16">
        <n-empty v-if="sessions.length === 0" :description="t('common.empty')" />
        <n-table v-else :bordered="false" :single-line="false" size="small">
          <thead>
            <tr>
              <th>{{ t("system.sessionId") }}</th>
              <th>{{ t("system.sessionName") }}</th>
              <th>{{ t("system.creatorPid") }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="s in sessions" :key="s.id">
              <td class="mono">{{ s.id }}</td>
              <td class="mono">{{ s.name }}</td>
              <td class="mono">{{ s.creatorPid }}</td>
            </tr>
          </tbody>
        </n-table>
      </n-card>
    </template>
  </div>
</template>

<style scoped>
.mt-16 {
  margin-top: 16px;
}
</style>
