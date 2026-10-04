import { defineStore } from "pinia";
import { ref } from "vue";

// 全局 UI 状态：语言（主题固定为浅色）
export const useAppStore = defineStore("app", () => {
  const locale = ref<"zh" | "en">("en");

  function setLocale(l: "zh" | "en") {
    locale.value = l;
  }

  return { locale, setLocale };
});
