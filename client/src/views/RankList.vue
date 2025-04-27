<template>
  <div class="ranklist">
    <div class="title">歌曲热榜</div>
    <song-list :songList="songList" @changeData="getAllSong"></song-list>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import SongList from "@/components/SongList.vue";
import { HttpManager } from "@/api";

const songList = ref([]);

// 获取所有歌曲
const getAllSong = async () => {
  try {
    const result = await HttpManager.getAllSong();
    if (result.success) {
      songList.value = result.data;
    }
  } catch (error) {
    console.error("获取歌曲列表失败:", error);
  }
};

onMounted(() => {
  getAllSong();
});
</script>

<style lang="scss" scoped>
@import "@/assets/css/var.scss";
@import "@/assets/css/global.scss";

.ranklist {
  padding: 20px;
  .title {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
    color: $color-black;
  }
}
</style>