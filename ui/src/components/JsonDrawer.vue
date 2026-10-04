<script setup lang="ts">
import { NDrawer, NDrawerContent, NSpin } from "naive-ui";

const props = defineProps<{
  show: boolean;
  title: string;
  loading?: boolean;
  content?: string | null;
}>();
const emit = defineEmits<{ (e: "update:show", v: boolean): void }>();
</script>

<template>
  <n-drawer
    :show="show"
    placement="right"
    :width="560"
    @update:show="(v: boolean) => emit('update:show', v)"
  >
    <n-drawer-content :title="title">
      <n-spin v-if="loading" :show="true">
        <div style="min-height: 140px" />
      </n-spin>
      <pre v-else class="json-block">{{ content || "—" }}</pre>
    </n-drawer-content>
  </n-drawer>
</template>

<style scoped>
.json-block {
  white-space: pre-wrap;
  word-break: break-all;
  font-family: "Cascadia Code", "Fira Code", Consolas, monospace;
  font-size: 12px;
  line-height: 1.6;
  margin: 0;
}
</style>
