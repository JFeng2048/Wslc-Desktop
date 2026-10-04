import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "system",
    component: () => import("@/views/SystemStatus.vue"),
  },
  {
    path: "/containers",
    name: "containers",
    component: () => import("@/views/ContainerList.vue"),
  },
  {
    path: "/images",
    name: "images",
    component: () => import("@/views/ImageList.vue"),
  },
  {
    path: "/volumes",
    name: "volumes",
    component: () => import("@/views/VolumeList.vue"),
  },
  {
    path: "/networks",
    name: "networks",
    component: () => import("@/views/NetworkList.vue"),
  },
  {
    path: "/terminal",
    name: "terminal",
    component: () => import("@/views/Terminal.vue"),
  },
];

export default createRouter({
  history: createWebHashHistory(),
  routes,
});
