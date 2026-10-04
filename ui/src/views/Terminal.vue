<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useMessage, NInput, NButton, NSpace, NTag } from "naive-ui";
import { Terminal as TerminalIcon, Play } from "lucide-vue-next";
import { runCommand } from "@/api/command";
import { toMessage } from "@/utils/error";
import type { CommandResult } from "@/types";

const { t } = useI18n();
const message = useMessage();

const command = ref("");
const running = ref(false);
const result = ref<CommandResult | null>(null);
const history = ref<string[]>([]);

const examples = ["images", "list -a", "network ls", "info", "version", "system prune"];

async function execute(raw?: string) {
  const cmd = (raw ?? command.value).trim();
  if (!cmd) return;
  const args = cmd.split(/\s+/);
  running.value = true;
  result.value = null;
  try {
    result.value = await runCommand(args);
    if (!history.value.includes(cmd)) history.value.unshift(cmd);
  } catch (e) {
    message.error(t("terminal.opFailed", { msg: toMessage(e) }));
  } finally {
    running.value = false;
  }
}

onMounted(() => {
  command.value = "images";
});
</script>

<template>
  <div class="page">
    <div class="page-header">
      <h2 class="page-title">{{ t("terminal.title") }}</h2>
    </div>
    <p class="page-tip">{{ t("terminal.note") }}</p>

    <div class="term-bar">
      <n-input
        v-model:value="command"
        :placeholder="t('terminal.placeholder')"
        clearable
        @keyup.enter="execute()"
      >
        <template #prefix><TerminalIcon :size="14" /></template>
      </n-input>
      <n-button type="primary" :loading="running" @click="execute()">
        <template #icon><Play :size="16" /></template>
        {{ t("terminal.run") }}
      </n-button>
    </div>

    <div class="examples">
      <span class="muted">{{ t("terminal.example") }}:</span>
      <n-tag
        v-for="ex in examples"
        :key="ex"
        class="ex-tag"
        size="small"
        round
        tertiary
        @click="execute(ex)"
      >
        wslc {{ ex }}
      </n-tag>
    </div>

    <div class="term-output">
      <div v-if="!result" class="muted term-placeholder">{{ t("terminal.output") }}</div>
      <template v-else>
        <div class="term-meta">
          <n-tag size="small" :type="result.exitCode === 0 ? 'success' : 'error'">
            exit {{ result.exitCode }}
          </n-tag>
        </div>
        <pre v-if="result.stdout" class="term-stdout">{{ result.stdout }}</pre>
        <pre v-if="result.stderr" class="term-stderr">{{ result.stderr }}</pre>
        <div v-if="!result.stdout && !result.stderr" class="muted">{{ t("terminal.noOutput") }}</div>
      </template>
    </div>
  </div>
</template>

<style scoped>
.term-bar {
  display: flex;
  gap: 8px;
  align-items: center;
}
.examples {
  margin: 12px 0;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
}
.ex-tag {
  cursor: pointer;
}
.term-output {
  margin-top: 12px;
  background: var(--card-color, #ffffff);
  border: 1px solid var(--border-color, #e5e7eb);
  border-radius: 8px;
  padding: 14px;
  min-height: 240px;
}
.term-stdout,
.term-stderr {
  white-space: pre-wrap;
  word-break: break-all;
  font-family: "Cascadia Code", "Fira Code", Consolas, monospace;
  font-size: 12.5px;
  line-height: 1.6;
  margin: 8px 0 0;
}
.term-stderr {
  color: #e57373;
}
.term-meta {
  margin-bottom: 4px;
}
.term-placeholder {
  padding: 20px 0;
}
</style>
