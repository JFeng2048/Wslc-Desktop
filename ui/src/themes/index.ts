import { lightTheme, type GlobalThemeOverrides } from "naive-ui";

// 品牌主色（青色，呼应终端 / 容器运维工具的气质）
export const brandColor = "#22d3ee";

// 浅色主题覆盖（仅保留浅色主题）
export const lightThemeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: "#0891b2",
    primaryColorHover: "#06b6d4",
    primaryColorPressed: "#0e7490",
    primaryColorSuppl: "#06b6d4",
    borderRadius: "8px",
    fontFamily:
      '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "PingFang SC", "Microsoft YaHei", sans-serif',
  },
};

export { lightTheme };
