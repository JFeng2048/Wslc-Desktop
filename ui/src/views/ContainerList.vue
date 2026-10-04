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
  NSelect,
  NSpace,
  NEmpty,
} from "naive-ui";
import type { SelectOption, DataTableColumn } from "naive-ui";
import {
  Play,
  Square,
  RotateCw,
  Zap,
  Eye,
  ScrollText,
  Activity,
  Trash2,
  RefreshCw,
  Boxes,
  Trash,
} from "lucide-vue-next";
import { storeToRefs } from "pinia";
import { useContainerStore } from "@/stores/container";
import type { Container } from "@/types";
import { toMessage } from "@/utils/error";
import JsonDrawer from "@/components/JsonDrawer.vue";

const { t } = useI18n();
const message = useMessage();
const dialog = useDialog();
const store = useContainerStore();
const { list, loading } = storeToRefs(store);

const keyword = ref("");
const stateFilter = ref("");

const stateType = (s: string) =>
  s === "running"
    ? "success"
    : s === "paused"
      ? "warning"
      : s === "exited"
        ? "error"
        : "default";

const stateOptions = computed<SelectOption[]>(() => {
  const set = new Set(list.value.map((c) => c.state).filter(Boolean));
  return [
    { label: t("common.all"), value: "" },
    ...[...set].map((v) => ({ label: v, value: v })),
  ];
});

const filtered = computed(() =>
  list.value.filter((c) => {
    const k = keyword.value.trim().toLowerCase();
    const okK =
      !k ||
      c.names.toLowerCase().includes(k) ||
      c.id.toLowerCase().includes(k) ||
      c.image.toLowerCase().includes(k);
    const okS = !stateFilter.value || c.state === stateFilter.value;
    return okK && okS;
  }),
);

const acting = ref<Record<string, string>>({});

async function run(fn: () => Promise<void>, successKey: string) {
  try {
    await fn();
    message.success(t(successKey));
    await store.fetch();
  } catch (e) {
    message.error(t("container.opFailed", { msg: toMessage(e) }));
  }
}

function confirmRemove(row: Container) {
  dialog.warning({
    title: t("common.delete"),
    content: t("container.confirmDeleteForce", { name: row.names || row.id }),
    positiveText: t("common.delete"),
    negativeText: t("common.cancel"),
    onPositiveClick: () =>
      run(() => store.remove(row.id, true), "container.deleteSuccess"),
  });
}

function confirmPrune() {
  dialog.warning({
    title: t("common.prune"),
    content: t("container.confirmPrune"),
    positiveText: t("common.prune"),
    negativeText: t("common.cancel"),
    onPositiveClick: () =>
      run(() => store.prune().then(() => {}), "container.pruneSuccess"),
  });
}

// 详情抽屉（inspect / logs / stats 复用）
const detail = ref<{ show: boolean; title: string; loading: boolean; content: string | null }>({
  show: false,
  title: "",
  loading: false,
  content: null,
});

async function openDetail(row: Container, kind: "inspect" | "logs" | "stats") {
  const titles = {
    inspect: t("container.inspectTitle"),
    logs: t("container.logsTitle"),
    stats: t("container.statsTitle"),
  } as const;
  detail.value = { show: true, title: `${titles[kind]} · ${row.names || row.id}`, loading: true, content: null };
  try {
    let content = "";
    if (kind === "inspect") content = await store.inspect(row.id);
    else if (kind === "logs") content = await store.logs(row.id, 200, true);
    else content = await store.stats();
    detail.value.content = content || t("common.empty");
  } catch (e) {
    detail.value.content = toMessage(e);
  } finally {
    detail.value.loading = false;
  }
}

const columns = computed<DataTableColumn<Container>[]>(() => [
  {
    title: t("container.name"),
    key: "names",
    minWidth: 160,
    render: (row: Container) => h("span", { class: "mono" }, row.names || row.id),
  },
  {
    title: t("container.image"),
    key: "image",
    minWidth: 160,
    render: (row: Container) => h("span", { class: "mono" }, row.image),
  },
  {
    title: t("container.state"),
    key: "state",
    width: 110,
    render: (row: Container) =>
      h(NTag, { type: stateType(row.state), size: "small" }, { default: () => row.state || t("state.unknown") }),
  },
  {
    title: t("container.ports"),
    key: "ports",
    minWidth: 140,
    render: (row: Container) => h("span", { class: "mono muted" }, row.ports || "—"),
  },
  {
    title: t("container.created"),
    key: "createdAt",
    width: 180,
    render: (row: Container) => h("span", { class: "muted" }, row.createdAt || "—"),
  },
  {
    title: t("common.actions"),
    key: "actions",
    width: 320,
    fixed: "right",
    render: (row: Container) => {
      const running = row.state === "running";
      return h(NSpace, { size: 4, wrap: false }, {
        default: () => [
          h(NButton, {
            size: "small", type: "primary", secondary: true, disabled: running,
            loading: acting.value[row.id] === "start",
            onClick: () => act(row, "start"),
          }, { default: () => t("common.start"), icon: () => h(Play, { size: 14 }) }),
          h(NButton, {
            size: "small", secondary: true, disabled: !running,
            onClick: () => act(row, "stop"),
          }, { default: () => t("common.stop"), icon: () => h(Square, { size: 14 }) }),
          h(NButton, {
            size: "small", secondary: true, loading: acting.value[row.id] === "restart",
            onClick: () => act(row, "restart"),
          }, { default: () => t("common.restart"), icon: () => h(RotateCw, { size: 14 }) }),
          h(NButton, {
            size: "small", secondary: true, disabled: running,
            onClick: () => act(row, "kill"),
          }, { default: () => t("container.kill"), icon: () => h(Zap, { size: 14 }) }),
          h(NButton, {
            size: "small", tertiary: true,
            onClick: () => openDetail(row, "inspect"),
          }, { default: () => t("common.inspect"), icon: () => h(Eye, { size: 14 }) }),
          h(NButton, {
            size: "small", tertiary: true,
            onClick: () => openDetail(row, "logs"),
          }, { default: () => t("common.logs"), icon: () => h(ScrollText, { size: 14 }) }),
          h(NButton, {
            size: "small", tertiary: true,
            onClick: () => openDetail(row, "stats"),
          }, { default: () => t("common.stats"), icon: () => h(Activity, { size: 14 }) }),
          h(NButton, {
            size: "small", type: "error", secondary: true,
            onClick: () => confirmRemove(row),
          }, { default: () => t("common.delete"), icon: () => h(Trash2, { size: 14 }) }),
        ],
      });
    },
  },
]);

async function act(row: Container, kind: "start" | "stop" | "restart" | "kill") {
  acting.value = { ...acting.value, [row.id]: kind };
  const map = {
    start: ["container.startSuccess", () => store.start(row.id)],
    stop: ["container.stopSuccess", () => store.stop(row.id)],
    restart: ["container.restartSuccess", () => store.restart(row.id)],
    kill: ["container.killSuccess", () => store.kill(row.id)],
  } as const;
  const [key, fn] = map[kind];
  await run(fn, key);
  const next = { ...acting.value };
  delete next[row.id];
  acting.value = next;
}

const rowKey = (row: Container) => row.id;

onMounted(store.fetch);
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h2 class="page-title">{{ t("container.title") }}</h2>
      <n-space>
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
      <n-input
        v-model:value="keyword"
        :placeholder="t('common.search')"
        clearable
        style="max-width: 280px"
      >
        <template #prefix><Boxes :size="14" /></template>
      </n-input>
      <n-select
        v-model:value="stateFilter"
        :options="stateOptions"
        :placeholder="t('container.filterState')"
        style="max-width: 180px"
      />
    </div>
    <n-data-table
      :columns="columns"
      :data="filtered"
      :loading="loading"
      :row-key="rowKey"
      :scroll-x="1100"
      size="small"
    />
    <n-empty
      v-if="!loading && filtered.length === 0"
      :description="t('container.none')"
      style="margin-top: 40px"
    />
    <JsonDrawer
      v-model:show="detail.show"
      :title="detail.title"
      :loading="detail.loading"
      :content="detail.content"
    />
  </div>
</template>
