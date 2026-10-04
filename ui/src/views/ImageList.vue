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
import { Trash2, RefreshCw, Image as ImageIcon, Download, Upload, Tag as TagIcon, Eye, Trash } from "lucide-vue-next";
import { storeToRefs } from "pinia";
import { useImageStore } from "@/stores/image";
import type { Image } from "@/types";
import { toMessage } from "@/utils/error";
import JsonDrawer from "@/components/JsonDrawer.vue";

const { t } = useI18n();
const message = useMessage();
const dialog = useDialog();
const store = useImageStore();
const { list, loading } = storeToRefs(store);

const keyword = ref("");
const filtered = computed(() =>
  list.value.filter((im) => {
    const k = keyword.value.trim().toLowerCase();
    if (!k) return true;
    return (
      (im.repository + ":" + im.tag).toLowerCase().includes(k) ||
      im.id.toLowerCase().includes(k)
    );
  }),
);

async function run(fn: () => Promise<void>, successKey: string) {
  try {
    await fn();
    message.success(t(successKey));
    await store.fetch();
  } catch (e) {
    message.error(t("image.opFailed", { msg: toMessage(e) }));
  }
}

function confirmRemove(row: Image) {
  dialog.warning({
    title: t("common.delete"),
    content: t("image.confirmDeleteForce", { repo: row.repository, tag: row.tag }),
    positiveText: t("common.delete"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => run(() => store.remove(row.id, true), "image.deleteSuccess"),
  });
}

function confirmPrune() {
  dialog.warning({
    title: t("common.prune"),
    content: t("image.confirmPrune"),
    positiveText: t("common.prune"),
    negativeText: t("common.cancel"),
    onPositiveClick: () => run(() => store.prune().then(() => {}), "image.pruneSuccess"),
  });
}

// inspect 抽屉
const detail = ref<{ show: boolean; title: string; loading: boolean; content: string | null }>({
  show: false,
  title: "",
  loading: false,
  content: null,
});
async function openInspect(row: Image) {
  detail.value = { show: true, title: `${t("common.inspect")} · ${row.repository}:${row.tag}`, loading: true, content: null };
  try {
    detail.value.content = (await store.inspect(row.id)) || t("common.empty");
  } catch (e) {
    detail.value.content = toMessage(e);
  } finally {
    detail.value.loading = false;
  }
}

// pull / push / tag 弹窗
const modal = ref<{
  type: "" | "pull" | "push" | "tag";
  show: boolean;
  ref: string;
  source: string;
  target: string;
}>({ type: "", show: false, ref: "", source: "", target: "" });

function openPull() {
  modal.value = { type: "pull", show: true, ref: "", source: "", target: "" };
}
function openPush() {
  modal.value = { type: "push", show: true, ref: "", source: "", target: "" };
}
function openTag(row: Image) {
  modal.value = {
    type: "tag",
    show: true,
    ref: "",
    source: `${row.repository}:${row.tag}`,
    target: "",
  };
}
async function submitModal() {
  const m = modal.value;
  try {
    if (m.type === "pull") await store.pull(m.ref);
    else if (m.type === "push") await store.push(m.ref);
    else if (m.type === "tag") await store.tag(m.source, m.target);
    message.success(t("common.success"));
    m.show = false;
    await store.fetch();
  } catch (e) {
    message.error(t("image.opFailed", { msg: toMessage(e) }));
  }
}

const columns = computed<DataTableColumn<Image>[]>(() => [
  {
    title: t("image.repository"),
    key: "repository",
    minWidth: 160,
    render: (row: Image) => h("span", { class: "mono" }, row.repository),
  },
  {
    title: t("image.tag"),
    key: "tag",
    width: 120,
    render: (row: Image) =>
      h(NTag, { size: "small", type: "info" }, { default: () => row.tag || "latest" }),
  },
  {
    title: t("image.id"),
    key: "id",
    minWidth: 160,
    render: (row: Image) => h("span", { class: "mono muted" }, row.id),
  },
  {
    title: t("image.size"),
    key: "size",
    width: 110,
    render: (row: Image) => h("span", { class: "muted" }, row.size || "—"),
  },
  {
    title: t("image.created"),
    key: "createdAt",
    width: 170,
    render: (row: Image) => h("span", { class: "muted" }, row.createdAt || "—"),
  },
  {
    title: t("common.actions"),
    key: "actions",
    width: 200,
    fixed: "right",
    render: (row: Image) =>
      h(NSpace, { size: 4, wrap: false }, {
        default: () => [
          h(NButton, {
            size: "small", tertiary: true,
            onClick: () => openInspect(row),
          }, { default: () => t("common.inspect"), icon: () => h(Eye, { size: 14 }) }),
          h(NButton, {
            size: "small", tertiary: true,
            onClick: () => openTag(row),
          }, { default: () => t("common.tag"), icon: () => h(TagIcon, { size: 14 }) }),
          h(NButton, {
            size: "small", type: "error", secondary: true,
            onClick: () => confirmRemove(row),
          }, { default: () => t("common.delete"), icon: () => h(Trash2, { size: 14 }) }),
        ],
      }),
  },
]);

const rowKey = (row: Image) => row.id;

onMounted(store.fetch);
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h2 class="page-title">{{ t("image.title") }}</h2>
      <n-space>
        <n-button secondary strong @click="openPull()">
          <template #icon><Download :size="16" /></template>
          {{ t("common.pull") }}
        </n-button>
        <n-button secondary strong @click="openPush()">
          <template #icon><Upload :size="16" /></template>
          {{ t("common.push") }}
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
      <n-input
        v-model:value="keyword"
        :placeholder="t('common.search')"
        clearable
        style="max-width: 280px"
      >
        <template #prefix><ImageIcon :size="14" /></template>
      </n-input>
    </div>
    <n-data-table
      :columns="columns"
      :data="filtered"
      :loading="loading"
      :row-key="rowKey"
      :scroll-x="820"
      size="small"
    />
    <n-empty
      v-if="!loading && filtered.length === 0"
      :description="t('image.none')"
      style="margin-top: 40px"
    />

    <JsonDrawer
      v-model:show="detail.show"
      :title="detail.title"
      :loading="detail.loading"
      :content="detail.content"
    />

    <n-modal
      v-model:show="modal.show"
      :title="modal.type === 'pull' ? t('common.pull') : modal.type === 'push' ? t('common.push') : t('common.tag')"
      preset="card"
      style="width: 460px"
    >
      <n-form v-if="modal.type !== 'tag'" label-placement="top">
        <n-form-item :label="t('common.imageRef')">
          <n-input v-model:value="modal.ref" :placeholder="t('image.refPlaceholder')" />
        </n-form-item>
      </n-form>
      <n-form v-else label-placement="top">
        <n-form-item :label="t('image.source')">
          <n-input v-model:value="modal.source" :placeholder="t('image.sourcePlaceholder')" />
        </n-form-item>
        <n-form-item :label="t('image.target')">
          <n-input v-model:value="modal.target" :placeholder="t('image.targetPlaceholder')" />
        </n-form-item>
      </n-form>
      <n-space justify="end">
        <n-button @click="modal.show = false">{{ t("common.cancel") }}</n-button>
        <n-button type="primary" @click="submitModal()">{{ t("common.confirm") }}</n-button>
      </n-space>
    </n-modal>
  </div>
</template>
