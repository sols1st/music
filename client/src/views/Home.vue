<template>
    <!--轮播图-->
    <el-carousel
        v-if="swiperList.length"
        class="swiper-container"
        type="card"
        height="18vw"
        :interval="4000"
    >
        <el-carousel-item
            class="item"
            v-for="(item, index) in swiperList"
            :key="index"
        >
            <img :src="HttpManager.attachImageUrl(item.pic)" />
        </el-carousel-item>
    </el-carousel>
    <!--热门歌单-->
    <play-list
        class="play-list-container"
        title="精选歌单"
        path="song-sheet-detail"
        :playList="songList"
    ></play-list>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";

import PlayList from "@/components/PlayList.vue";
import { NavName } from "@/enums";
import { HttpManager } from "@/api";
import mixin from "@/mixins/mixin";

const songList = ref([]); // 歌单列表
const singerList = ref([]); // 歌手列表
const swiperList = ref([]); // 轮播图 每次都在进行查询
const { changeIndex } = mixin();
try {
    HttpManager.getBannerList().then((res) => {
        swiperList.value = (res as ResponseBody).data.sort();
    });

    HttpManager.getSongList().then((res) => {
        songList.value = (res as ResponseBody).data.sort().slice(0, 10);
    });

    HttpManager.getAllSinger().then((res) => {
        singerList.value = (res as ResponseBody).data.sort().slice(0, 10);
    });

    onMounted(() => {
        changeIndex(NavName.Home);
    });
} catch (error) {
    console.error(error);
}
</script>

<style lang="scss" scoped>
@import "@/assets/css/var.scss";

.swiper-container {
    width: 90%;
    margin: auto;
    padding-top: 20px;
    img {
        width: 100%;
        border-radius: 10px;
    }
}

.swiper-container:deep(
        .el-carousel__indicators.el-carousel__indicators--outside
    ) {
    display: flex;
    justify-content: center;
    transform: none;
}

.el-slider__runway {
    background-color: $color-blue;
}
</style>
