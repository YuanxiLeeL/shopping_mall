import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createPinia } from "pinia";
const pinia = createPinia();

createApp(App).use(Antd).use(router).use(pinia).mount("#app");

import Antd from "ant-design-vue";
import "ant-design-vue/dist/reset.css";

const app = createApp(App);
