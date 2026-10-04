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
import { Trash2, RefreshCw, Network as NetworkIcon, Eye, Trash, Plus } from "lucide-vue-next";
import { storeToRefs } from "pinia";
import { useNetworkStore } from "@/stores/network";
import type { Network } from "@/types";
import { toMessage } from "@/utils/error";
import JsonDrawer from "@/components/JsonDrawer.vue";

const { t } = useI18n();
const message = useMessage();
const dialog = useDialog();
const store = useNetworkStore();
const { list, loading } = storeToRefs(store);

const keyword = ref("");
const filtered = computed(() =>
  list.value.filter((n) => {
    const k = keyword.value.trim().toLowerCase();
    return !k || n.name.toLowerCase().includes(k) || n.id.toLowerCase().includes(k);
  }),
);

async function run(fn: () => Promise<void>, successKey: string) {
  try {
    await fn();
    message.success(t(successKey));
    await store.fetch();
  } catch (e) {
    message.error(t("network.opFailed", { msg: toMessage(e) }));
  }
}

function confirmRemove(row: Network) {
  dialog.warning({
    title: t("common.delete"),
    content: t("network.confirmDelete", { name: row.name }),
    positiveText: t("common.delete"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => run(() => store.remove(row.id), "network.deleteSuccess"),
  });
}
function confirmPrune() {
  dialog.warning({
    title: t("common.prune"),
    content: t("network.confirmPrune"),
    positiveText: t("common.prune"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => run(() => store.prune().then(() => {}), "network.pruneSuccess"),
  });
}

const detail = ref<{ show: boolean; title: string; loading: boolean; content: string | null }>({
  show: false, title: "", loading: false, content: null,
});
async function openInspect(row: Network) {
  detail.value = { show: true, title: `${t("common.inspect")} · ${row.name}`, loading: true, content: null };
  try {
    detail.value.content = (await store.inspect(row.id)) || t("common.empty");
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
    message.success(t("network.createSuccess"));
    showCreate.value = false;
    createName.value = "";
    await store.fetch();
  } catch (e) {
    message.error(t("network.opFailed", { msg: toMessage(e) }));
  }
}

const yn = (v: boolean) => (v ? t("common.yes") : t("common.no"));

const columns = computed<DataTableColumn<Network>[]>(() => [
  { title: t("network.name"), key: "name", minWidth: 160, render: (row: Network) => h("span", { class: "mono" }, row.name) },
  { title: t("network.id"), key: "id", minWidth: 140, render: (row: Network) => h("span", { class: "mono muted" }, row.id) },
  { title: t("network.driver"), key: "driver", width: 120, render: (row: Network) => h(NTag, { size: "small" }, { default: () => row.driver || "—" }) },
  { title: t("network.scope"), key: "scope", width: 100, render: (row: Network) => h("span", { class: "muted" }, row.scope || "—") },
  { title: t("network.ipv4"), key: "ipv4", width: 80, render: (row: Network) => h("span", { class: "muted" }, yn(row.ipv4)) },
  { title: t("network.ipv6"), key: "ipv6", width: 80, render: (row: Network) => h("span", { class: "muted" }, yn(row.ipv6)) },
  { title: t("network.internal"), key: "internal", width: 90, render: (row: Network) => h("span", { class: "muted" }, yn(row.internal)) },
  {
    title: t("common.actions"),
    key: "actions",
    width: 160,
    fixed: "right",
    render: (row: Network) =>
      h(NSpace, { size: 4, wrap: false }, {
        default: () => [
          h(NButton, { size: "small", tertiary: true, onClick: () => openInspect(row) }, { default: () => t("common.inspect"), icon: () => h(Eye, { size: 14 }) }),
          h(NButton, { size: "small", type: "error", secondary: true, onClick: () => confirmRemove(row) }, { default: () => t("common.delete"), icon: () => h(Trash2, { size: 14 }) }),
        ],
      }),
  },
]);

const rowKey = (row: Network) => row.id;
onMounted(store.fetch);
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h2 class="page-title">{{ t("network.title") }}</h2>
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
        <template #prefix><NetworkIcon :size="14" /></template>
      </n-input>
    </div>
    <n-data-table :columns="columns" :data="filtered" :loading="loading" :row-key="rowKey" :scroll-x="900" size="small" />
    <n-empty v-if="!loading && filtered.length === 0" :description="t('network.none')" style="margin-top: 40px" />

    <JsonDrawer v-model:show="detail.show" :title="detail.title" :loading="detail.loading" :content="detail.content" />

    <n-modal v-model:show="showCreate" :title="t('common.create')" preset="card" style="width: 420px">
      <n-form label-placement="top">
        <n-form-item :label="t('network.name')">
          <n-input v-model:value="createName" :placeholder="t('network.createPlaceholder')" />
        </n-form-item>
      </n-form>
      <n-space justify="end">
        <n-button @click="showCreate = false">{{ t("common.cancel") }}</n-button>
        <n-button type="primary" :disabled="!createName.trim()" @click="submitCreate()">{{ t("common.create") }}</n-button>
      </n-space>
    </n-modal>
  </div>
</template>
