import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import "./assets/styles/global.css";

const app = createApp(App).use(createPinia()).use(router).use(i18n);

// 全局错误兜底：避免单个视图在渲染/切换时抛错导致整页白屏
app.config.errorHandler = (err, _instance, info) => {
  console.error("[app error]", info, err);
};

app.mount("#app");
