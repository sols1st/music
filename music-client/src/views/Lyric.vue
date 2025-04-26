<template>
  <div class="song-container">
    <el-image class="song-pic" fit="contain" :src="attachImageUrl(songPic)" />
    <ul class="song-info">
      <li class="title">歌手：</li>
      <li class="info">{{ singerName }}</li>
      <li class="title">歌曲：</li>
      <li class="info">{{ songTitle }}</li>
    </ul>
  </div>

  <div class="container">
    <div class="lyric-container">
      <div class="song-lyric">
        <transition-group name="lyric-fade">
          <ul :style="{ top: lrcTop }" class="has-lyric" v-if="lyricArr.length" key="has-lyric">
            <li v-for="(item, index) in lyricArr" :key="index">
              {{ item[1] }}
            </li>
          </ul>
          <div v-else class="no-lyric" key="no-lyric">
            <span>暂无歌词</span>
          </div>
        </transition-group>
      </div>
    <comment :playId="songId" :type="0"></comment>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref, watch } from "vue";
import { useStore } from "vuex";
import Comment from "@/components/Comment.vue";
import { parseLyric } from "@/utils";
import { HttpManager } from "@/api";

export default defineComponent({
  components: {
    Comment,
  },
  setup() {
    const store = useStore();

    const lrcTop = ref("80px"); // 歌词滑动
    const lyricArr = ref([]); // 当前歌曲的歌词
    const songId = computed(() => store.getters.songId); // 歌曲ID
    const lyric = computed(() => store.getters.lyric); // 歌词
    const currentPlayList = computed(() => store.getters.currentPlayList); // 存放的音乐
    const currentPlayIndex = computed(() => store.getters.currentPlayIndex); // 当前歌曲在歌曲列表的位置
    const curTime = computed(() => store.getters.curTime);
    const songTitle = computed(() => store.getters.songTitle); // 歌名
    const singerName = computed(() => store.getters.singerName); // 歌手名
    const songPic = computed(() => store.getters.songPic); // 歌曲图片
    watch(songId, () => {
      lyricArr.value = parseLyric(currentPlayList.value[currentPlayIndex.value].lyric);
    });
    // 处理歌词位置及颜色
    watch(curTime, () => {
      if (lyricArr.value.length !== 0) {
        for (let i = 0; i < lyricArr.value.length; i++) {
          if (curTime.value >= lyricArr.value[i][0]) {
            for (let j = 0; j < lyricArr.value.length; j++) {
              (document.querySelectorAll(".has-lyric li") as NodeListOf<HTMLElement>)[j].style.color = "#aab5bd";
              (document.querySelectorAll(".has-lyric li") as NodeListOf<HTMLElement>)[j].style.fontSize = "28px";
              (document.querySelectorAll(".has-lyric li") as NodeListOf<HTMLElement>)[j].style.fontWeight = "bold";
            }
            if (i >= 0) {
              lrcTop.value = -i * 30 + 50 + "px";
              (document.querySelectorAll(".has-lyric li") as NodeListOf<HTMLElement>)[i].style.color = "#000";
              (document.querySelectorAll(".has-lyric li") as NodeListOf<HTMLElement>)[i].style.fontSize = "30px";
            }
          }
        }
      }
    });

    lyricArr.value = lyric.value ? parseLyric(lyric.value) : [];

    return {
      songPic,
      singerName,
      songTitle,
      lrcTop,
      lyricArr,
      songId,
      attachImageUrl: HttpManager.attachImageUrl,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "@/assets/css/var.scss";

.song-container {
  position: fixed;
  top: 120px;
  left: 50px;
  display: flex;
  flex-direction: column;

  .song-pic {
    height: 300px;
    border: 4px solid white;
    border-radius: 15px;
  }

  .song-info {
    width: 300px;
    .info{
      width: 100%;
      line-height: 50px;
      font-size: 40px;
    }
    .title{
      margin-top: 20px;
      width: 100%;
      line-height: 30px;
      color: $color-grey;
      font-size: 30px;
    }
  }
}

.lyric-container {
  font-family: $font-family;
  height: content-max;
  .song-lyric {
    position: relative;
    height: 100vh;
    padding: 30px 0;
    overflow: auto;
    border-radius: 15px;
    display: flex;
    justify-content: center;
    .has-lyric {
      position: absolute;
      transition: all 1s;
      li {
        width: 100%;
        height: 40px;
        text-align: center;
        font-size: 28px;
        font-weight: bold;
        line-height: 40px;
        color: $color-grey;
      }
    }
    .no-lyric {
      position: absolute;
      margin: 100px 0;

      span {
        font-size: 18px;
        text-align: center;
      }
    }
  }
}

.lyric-fade-enter,
.lyric-fade-leave-to {
  transform: translateX(30px);
  opacity: 0;
}

.lyric-fade-enter-active,
.lyric-fade-leave-active {
  transition: all 0.3s ease;
}

@media screen and (min-width: $sm) {
  .container {
    padding-top: 30px;
  }
  .lyric-container {
    margin: 0 150px 0px 400px;
  }
}

@media screen and (max-width: $sm) {
  .container {
    padding: 20px;
  }
  .song-container {
    display: none;
  }
}
</style>
