import { createApp } from "vue";
import ElementPlus from "element-plus";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "element-plus/dist/index.css";
import "./assets/css/index.scss";
import { setupSvgIcon } from "./assets/icons/index.js";

import { Store } from "vuex";

declare module "@vue/runtime-core" {
  interface State {
    count: number;
  }

  interface ComponentCustomProperties {
    $store: Store<State>;
  }
}

const app = createApp(App);
setupSvgIcon(app);

app.use(store)
   .use(router)
   .use(ElementPlus)
   .mount("#app");
