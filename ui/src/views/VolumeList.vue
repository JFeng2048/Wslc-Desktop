<script setup lang="ts">
import { computed, h, onMounted, ref } from "vue";
import { useI18n } from "vue-i18n";
import {
  useMessage,
  useDialog,
  NButton,
  NTag,
  NDataTable,
  NInput,
  NEmpty,
  NModal,
  NSpace,
  NForm,
  NFormItem,
} from "naive-ui";
import type { DataTableColumn } from "naive-ui";
import { Trash2, RefreshCw, Database, Eye, Trash, Plus } from "lucide-vue-next";
import { storeToRefs } from "pinia";
import { useVolumeStore } from "@/stores/volume";
import type { Volume } from "@/types";
import { toMessage } from "@/utils/error";
import JsonDrawer from "@/components/JsonDrawer.vue";

const { t } = useI18n();
const message = useMessage();
const dialog = useDialog();
const store = useVolumeStore();
const { list, loading } = storeToRefs(store);

const keyword = ref("");
const filtered = computed(() =>
  list.value.filter((v) => {
    const k = keyword.value.trim().toLowerCase();
    return !k || v.name.toLowerCase().includes(k) || (v.driver || "").toLowerCase().includes(k);
  }),
);

async function run(fn: () => Promise<void>, successKey: string) {
  try {
    await fn();
    message.success(t(successKey));
    await store.fetch();
  } catch (e) {
    message.error(t("volume.opFailed", { msg: toMessage(e) }));
  }
}

function confirmRemove(row: Volume) {
  dialog.warning({
    title: t("common.delete"),
    content: t("volume.confirmDelete", { name: row.name }),
    positiveText: t("common.delete"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => run(() => store.remove(row.name), "volume.deleteSuccess"),
  });
}
function confirmPrune() {
  dialog.warning({
    title: t("common.prune"),
    content: t("volume.confirmPrune"),
    positiveText: t("common.prune"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => run(() => store.prune().then(() => {}), "volume.pruneSuccess"),
  });
}

const detail = ref<{ show: boolean; title: string; loading: boolean; content: string | null }>({
  show: false, title: "", loading: false, content: null,
});
async function openInspect(row: Volume) {
  detail.value = { show: true, title: `${t("common.inspect")} · ${row.name}`, loading: true, content: null };
  try {
    detail.value.content = (await store.inspect(row.name)) || t("common.empty");
  } catch (e) {
    detail.value.content = toMessage(e);
  } finally {
    detail.value.loading = false;
  }
}

const createName = ref("");
const showCreate = ref(false);
async function submitCreate() {
  try {
    await store.create(createName.value.trim());
    message.success(t("volume.createSuccess"));
    showCreate.value = false;
    createName.value = "";
    await store.fetch();
  } catch (e) {
    message.error(t("volume.opFailed", { msg: toMessage(e) }));
  }
}

const columns = computed<DataTableColumn<Volume>[]>(() => [
  {
    title: t("volume.name"),
    key: "name",
    minWidth: 160,
    render: (row: Volume) => h("span", { class: "mono" }, row.name),
  },
  {
    title: t("volume.driver"),
    key: "driver",
    width: 120,
    render: (row: Volume) => h(NTag, { size: "small" }, { default: () => row.driver || "—" }),
  },
  { title: t("volume.scope"), key: "scope", width: 110, render: (row: Volume) => h("span", { class: "muted" }, row.scope || "—") },
  { title: t("volume.size"), key: "size", width: 100, render: (row: Volume) => h("span", { class: "muted" }, row.size || "—") },
  { title: t("volume.status"), key: "status", width: 100, render: (row: Volume) => h("span", { class: "muted" }, row.status || "—") },
  {
    title: t("common.actions"),
    key: "actions",
    width: 160,
    fixed: "right",
    render: (row: Volume) =>
      h(NSpace, { size: 4, wrap: false }, {
        default: () => [
          h(NButton, { size: "small", tertiary: true, onClick: () => openInspect(row) }, { default: () => t("common.inspect"), icon: () => h(Eye, { size: 14 }) }),
          h(NButton, { size: "small", type: "error", secondary: true, onClick: () => confirmRemove(row) }, { default: () => t("common.delete"), icon: () => h(Trash2, { size: 14 }) }),
        ],
      }),
  },
]);

const rowKey = (row: Volume) => row.name;
onMounted(store.fetch);
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h2 class="page-title">{{ t("volume.title") }}</h2>
      <n-space>
        <n-button secondary strong @click="showCreate = true">
          <template #icon><Plus :size="16" /></template>
          {{ t("common.create") }}
        </n-button>
        <n-button :loading="loading" secondary strong @click="confirmPrune()">
          <template #icon><Trash :size="16" /></template>
          {{ t("common.prune") }}
        </n-button>
        <n-button :loading="loading" secondary strong @click="store.fetch()">
          <template #icon><RefreshCw :size="16" /></template>
          {{ t("common.refresh") }}
        </n-button>
      </n-space>
    </div>
    <div class="toolbar">
      <n-input v-model:value="keyword" :placeholder="t('common.search')" clearable style="max-width: 280px">
        <template #prefix><Database :size="14" /></template>
      </n-input>
    </div>
    <n-data-table :columns="columns" :data="filtered" :loading="loading" :row-key="rowKey" :scroll-x="760" size="small" />
    <n-empty v-if="!loading && filtered.length === 0" :description="t('volume.none')" style="margin-top: 40px" />

    <JsonDrawer v-model:show="detail.show" :title="detail.title" :loading="detail.loading" :content="detail.content" />

    <n-modal v-model:show="showCreate" :title="t('common.create')" preset="card" style="width: 420px">
      <n-form label-placement="top">
        <n-form-item :label="t('volume.name')">
          <n-input v-model:value="createName" :placeholder="t('volume.createPlaceholder')" />
        </n-form-item>
      </n-form>
      <n-space justify="end">
        <n-button @click="showCreate = false">{{ t("common.cancel") }}</n-button>
        <n-button type="primary" :disabled="!createName.trim()" @click="submitCreate()">{{ t("common.create") }}</n-button>
      </n-space>
    </n-modal>
  </div>
</template>
