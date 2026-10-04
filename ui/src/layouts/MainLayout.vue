<script setup lang="ts">
import { computed, h, type Component } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import {
  NLayout,
  NLayoutSider,
  NLayoutHeader,
  NLayoutContent,
  NMenu,
  NButton,
  NIcon,
  NDropdown,
  NText,
} from "naive-ui";
import { Cpu, Boxes, Image as ImageIcon, Languages, Database, Network as NetworkIcon, Terminal as TerminalIcon } from "lucide-vue-next";
import { useAppStore } from "@/stores/app";

const { t } = useI18n();
const router = useRouter();
const route = useRoute();
const app = useAppStore();

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = computed(() => [
  { label: t("nav.system"), key: "/", icon: renderIcon(Cpu) },
  { label: t("nav.containers"), key: "/containers", icon: renderIcon(Boxes) },
  { label: t("nav.images"), key: "/images", icon: renderIcon(ImageIcon) },
  { label: t("nav.volumes"), key: "/volumes", icon: renderIcon(Database) },
  { label: t("nav.networks"), key: "/networks", icon: renderIcon(NetworkIcon) },
  { label: t("nav.terminal"), key: "/terminal", icon: renderIcon(TerminalIcon) },
]);

const activeKey = computed(() => route.path);

// 已是当前页则忽略，避免重复导航报错
function handleSelect(key: string) {
  if (key === route.path) return;
  router.push(key).catch(() => {});
}

const langOptions = [
  { label: "中文", key: "zh" },
  { label: "English", key: "en" },
];
function handleLang(key: string) {
  app.setLocale(key as "zh" | "en");
}
</script>

<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider bordered :width="220" :collapsed-width="64" show-trigger="bar">
      <div class="brand">
        <div class="brand-logo">W</div>
        <div class="brand-text">
          <div class="brand-title">{{ t("app.title") }}</div>
          <div class="brand-sub">{{ t("app.subtitle") }}</div>
        </div>
      </div>
      <n-menu :value="activeKey" :options="menuOptions" @update:value="handleSelect" />
    </n-layout-sider>
    <n-layout>
      <n-layout-header bordered class="app-header">
        <n-text strong style="font-size: 16px">{{ t("app.title") }}</n-text>
        <div class="header-actions">
          <n-dropdown :options="langOptions" @select="handleLang">
            <n-button quaternary circle>
              <template #icon><n-icon><Languages /></n-icon></template>
            </n-button>
          </n-dropdown>
        </div>
      </n-layout-header>
      <n-layout-content class="app-content">
        <router-view v-slot="{ Component }">
          <keep-alive>
            <component :is="Component" />
          </keep-alive>
        </router-view>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>
