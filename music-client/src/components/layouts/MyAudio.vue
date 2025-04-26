<template>
  <audio 
    :src="attachImageUrl(songUrl)" 
    controls="controls" 
    :ref="player" 
    preload="true" 
    @canplay="canplay" 
    @timeupdate="timeupdate" 
    @ended="ended"
    @play="handlePlay"
    @pause="handlePause"
  >
    <!--（1）属性：controls，preload（2）事件：canplay，timeupdate，ended（3）方法：play()，pause() -->
    <!--controls：向用户显示音频控件（播放/暂停/进度条/音量）-->
    <!--preload：属性规定是否在页面加载后载入音频-->
    <!--canplay：当音频/视频处于加载过程中时，会发生的事件-->
    <!--timeupdate：当目前的播放位置已更改时-->
    <!--ended：当目前的播放列表已结束时-->
  </audio>
</template>

<script lang="ts">
import { defineComponent, ref, getCurrentInstance, computed, watch, onMounted } from "vue";
import { useStore } from "vuex";
import { HttpManager } from "@/api";

export default defineComponent({
  setup() {
    const { proxy } = getCurrentInstance();
    const store = useStore();
    const divRef = ref<HTMLAudioElement>();
    const player = (el) => {
      divRef.value = el;
    };

    const muted = ref(true);
    const hasUserInteraction = ref(false);

    // 监听用户交互
    const setupUserInteraction = () => {
      const handleUserInteraction = () => {
        hasUserInteraction.value = true;
        // 移除事件监听器
        document.removeEventListener('click', handleUserInteraction);
        document.removeEventListener('touchstart', handleUserInteraction);
      };

      document.addEventListener('click', handleUserInteraction);
      document.addEventListener('touchstart', handleUserInteraction);
    };

    onMounted(() => {
      setupUserInteraction();
    });

    const songUrl = computed(() => store.getters.songUrl);
    const isPlay = computed(() => store.getters.isPlay);
    const volume = computed(() => store.getters.volume);
    const changeTime = computed(() => store.getters.changeTime);
    const autoNext = computed(() => store.getters.autoNext);

    // 监听播放还是暂停
    watch(isPlay, () => togglePlay());
    // 跳到指定时刻播放
    watch(changeTime, () => (divRef.value.currentTime = changeTime.value));
    watch(volume, (value) => (divRef.value.volume = value));

    // 开始 / 暂停
    async function togglePlay() {
      if (!hasUserInteraction.value) {
        // 如果用户还没有交互，先静音播放
        divRef.value.muted = true;
        try {
          await divRef.value.play();
          // 播放成功后取消静音
          divRef.value.muted = false;
        } catch (error) {
          console.error('自动播放失败，需要用户交互。', error);
        }
      } else {
        // 用户已经交互过，直接播放
        isPlay.value ? divRef.value.play() : divRef.value.pause();
      }
    }

    // 获取歌曲链接后准备播放
    function canplay() {
      proxy.$store.commit("setDuration", divRef.value.duration);
      if (muted.value) {
        divRef.value.muted = false;
        muted.value = false;
      }
      proxy.$store.commit("setIsPlay", true);
    }

    // 音乐播放时记录音乐的播放位置
    function timeupdate() {
      proxy.$store.commit("setCurTime", divRef.value.currentTime);
    }

    // 音乐播放结束时触发
    function ended() {
      proxy.$store.commit("setIsPlay", false);
      proxy.$store.commit("setCurTime", 0);
      proxy.$store.commit("setAutoNext", !autoNext.value);
    }

    // 处理播放事件
    function handlePlay() {
      hasUserInteraction.value = true;
    }

    // 处理暂停事件
    function handlePause() {
      // 可以在这里添加暂停时的逻辑
    }

    return {
      songUrl,
      player,
      canplay,
      timeupdate,
      ended,
      muted,
      handlePlay,
      handlePause,
      attachImageUrl: HttpManager.attachImageUrl,
    };
  },
});
</script>

<style scoped>
audio {
  display: none;
}
</style>
