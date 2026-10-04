<script setup lang="ts">
import { computed, watch } from "vue";
import {
  NConfigProvider,
  NMessageProvider,
  NDialogProvider,
  NGlobalStyle,
  lightTheme,
  zhCN,
  enUS,
  dateZhCN,
  dateEnUS,
} from "naive-ui";
import { useAppStore } from "@/stores/app";
import { lightThemeOverrides } from "@/themes";
import i18n from "@/i18n";
import MainLayout from "@/layouts/MainLayout.vue";

const app = useAppStore();

// 仅保留浅色主题
const theme = computed(() => lightTheme);
const themeOverrides = computed(() => lightThemeOverrides);
const naiveLocale = computed(() => (app.locale === "zh" ? zhCN : enUS));
const naiveDateLocale = computed(() =>
  app.locale === "zh" ? dateZhCN : dateEnUS,
);

// 语言切换同步给 vue-i18n
watch(
  () => app.locale,
  (l) => {
    i18n.global.locale.value = l;
  },
  { immediate: true },
);
</script>

<template>
  <n-config-provider
    :theme="theme"
    :theme-overrides="themeOverrides"
    :locale="naiveLocale"
    :date-locale="naiveDateLocale"
  >
    <n-message-provider>
      <n-dialog-provider>
        <n-global-style />
        <main-layout />
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>
