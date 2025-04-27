<template>
  <el-container class="container">
    <el-header>
      <my-header></my-header>
    </el-header>
    <el-main>
      <router-view />
      <current-play></current-play>
      <play-bar></play-bar>
      <scroll-top></scroll-top>
      <my-audio></my-audio>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import { getCurrentInstance } from "vue";
import MyHeader from "@/components/layouts/MyHeader.vue";
import CurrentPlay from "@/components/layouts/CurrentPlay.vue";
import PlayBar from "@/components/layouts/PlayBar.vue";
import ScrollTop from "@/components/layouts/ScrollTop.vue";
import MyAudio from "@/components/layouts/MyAudio.vue";

const { proxy } = getCurrentInstance();

if (sessionStorage.getItem("dataStore")) {
  proxy.$store.replaceState(Object.assign({}, proxy.$store.state, JSON.parse(sessionStorage.getItem("dataStore"))));
}

window.addEventListener("beforeunload", () => {
  sessionStorage.setItem("dataStore", JSON.stringify(proxy.$store.state));
});
</script>

<style lang="scss" scoped>
@import "@/assets/css/var.scss";
@import "@/assets/css/global.scss";

.container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.el-container {
  min-height: calc(100% - 60px);
  display: flex;
  flex-direction: column;
}

.el-header {
  padding: 0;
  flex-shrink: 0;
}

.el-main {
  padding-left: 0;
  padding-right: 0;
  flex: 1;
  overflow-x: hidden;
  position: relative;
}
</style>
