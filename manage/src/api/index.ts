import {deletes, get, getBaseURL, post} from './request'

const HttpManager = {
    // 获取图片信息
    attachImageUrl: (url) => url ? `${getBaseURL()}${url}`: "https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png",
    // =======================> 管理员 API 完成
    // 是否登录成功
    getLoginStatus: ({username, password}) => post(`admin/login/status`, {username, password}),

    // =======================> 用户 API 完成
    // 返回所有用户
    getAllUser: () => get(`user/all`),
    // 返回指定ID的用户
    getUserOfId: (id) => get(`user/detail?id=${id}`),
    // 删除用户
    deleteUser: (id) => get(`user/delete?id=${id}`),
    // 添加用户
    addUser: (user) => post(`user/add`, user),
    // 用户登录
    loginStatus: (user) => post(`user/login/status`, user),
    // 邮箱登录
    loginEmailStatus: (user) => post(`user/email/status`, user),
    // 重置密码
    resetPassword: (user) => post(`user/resetPassword`, user),
    // 更新用户信息
    updateUserMsg: (user) => post(`user/update`, user),
    // 更新密码
    updatePassword: (user) => post(`user/updatePassword`, user),
    // 更新头像
    updateUserAvatar: (id) => `${getBaseURL()}/user/avatar/update?id=${id}`,

    // =======================> 收藏列表 API 完成
    // 返回的指定用户ID收藏列表
    getCollectionOfUser: (userId) => get(`collection/detail?userId=${userId}`),
    // 删除收藏的歌曲
    deleteCollection: (userId, songId) => deletes(`collection/delete?userId=${userId}&&songId=${songId}`),
    // 添加收藏
    addCollection: (collection) => post(`collection/add`, collection),
    // 获取收藏状态
    getCollectionStatus: (collection) => post(`collection/status`, collection),

    // =======================> 评论列表 API 完成
    // 获得指定歌曲ID的评论列表
    getCommentOfSongId: (songId) => get(`comment/song/detail?songId=${songId}`),
    // 获得指定歌单ID的评论列表
    getCommentOfSongListId: (songListId) => get(`comment/songList/detail?songListId=${songListId}`),
    // 删除评论
    deleteComment: (id) => get(`comment/delete?id=${id}`),
    // 添加评论
    addComment: (comment) => post(`comment/add`, comment),
    // 更新评论
    updateCommentMsg: (comment) => post(`comment/like`, comment),

    // =======================> 歌手 API 完成
    // 返回所有歌手
    getAllSinger: () => get(`singer/all`),
    // 添加歌手
    setSinger: ({name, sex, birth, location, introduction}) => post(`singer/add`, {
        name,
        sex,
        birth,
        location,
        introduction
    }),
    // 更新歌手信息
    updateSingerMsg: ({id, name, sex, birth, location, introduction}) => post(`singer/update`, {
        id,
        name,
        sex,
        birth,
        location,
        introduction
    }),
    // 删除歌手
    deleteSinger: (id) => deletes(`singer/delete?id=${id}`),
    // 根据名字获取歌手
    getSingerOfName: (name) => get(`singer/name/detail?name=${name}`),
    // 根据性别获取歌手
    getSingerOfSex: (sex) => get(`singer/sex/detail?sex=${sex}`),
    // 更新歌手头像
    updateSingerPic: (id) => `${getBaseURL()}/singer/avatar/update?id=${id}`,

    // =======================> 歌曲 API  完成
    // 返回所有歌曲
    getAllSong: () => get(`song/all`),
    // 返回指定歌手ID的歌曲
    getSongOfSingerId: (singerId) => get(`song/singer/detail?singerId=${singerId}`),
    // 返回的指定歌曲ID的歌曲
    getSongOfId: (id) => get(`song/detail?id=${id}`),
    // 返回指定歌手名的歌曲
    getSongOfSingerName: (name) => get(`song/singerName/detail?name=${name}`),
    // 更新歌曲信息
    updateSongMsg: ({id, singerId, name, introduction, lyric}) => post(`song/update`, {
        id,
        singerId,
        name,
        introduction,
        lyric
    }),
    // 更新歌曲URL
    updateSongUrl: (id) => `${getBaseURL()}/song/url/update?id=${id}`,
    // 更新歌曲图片
    updateSongImg: (id) => `${getBaseURL()}/song/img/update?id=${id}`,
    // 更新歌曲歌词
    updateSongLrc: (id) => `${getBaseURL()}/song/lrc/update?id=${id}`,
    // 删除歌曲
    deleteSong: (id) => deletes(`song/delete?id=${id}`),
    // 添加歌曲
    addSong: (song) => post(`song/add`, song),
    // 搜索歌曲
    searchSongs: (keywords) => get(`song/search?keywords=${keywords}`),

    // =======================> 歌单 API 完成
    // 添加歌单
    setSongList: ({title, introduction, style}) => post(`songList/add`, {title, introduction, style}),
    // 获取全部歌单
    getSongList: () => get(`songList/all`),
    // 更新歌单信息
    updateSongListMsg: ({id, title, introduction, style}) => post(`songList/update`, {id, title, introduction, style}),
    // 删除歌单
    deleteSongList: (id) => get(`songList/delete?id=${id}`),
    // 根据标题获取歌单
    getSongListOfTitle: (title) => get(`songList/likeTitle/detail?title=${title}`),
    // 根据风格获取歌单
    getSongListOfStyle: (style) => get(`songList/style/detail?style=${style}`),
    // 更新歌单图片
    updateSongListPic: (id) => `${getBaseURL()}/songList/img/update?id=${id}`,

    // =======================> 歌单歌曲 API 完成
    // 给歌单添加歌曲
    setListSong: ({songId,songListId}) => post(`listSong/add`, {songId,songListId}),
    // 返回歌单里指定歌单ID的歌曲
    getListSongOfSongId: (songListId) => get(`listSong/detail?songListId=${songListId}`),
    // 删除歌单里的歌曲
    deleteListSong: (songId) => get(`listSong/delete?songId=${songId}`),
    // 更新歌单歌曲信息
    updateListSongMsg: (listSong) => post(`listSong/update`, listSong),
    // 获取歌单歌曲Excel
    getListSongExcel: () => get(`listSong/excle`),

    // =======================> 轮播图 API 完成
    // 获取所有轮播图
    getAllBanner: () => get(`banner/getAllBanner`),

    // =======================> 文件下载 API 完成
    // 下载文件
    downloadFile: (fileName) => get(`download/${fileName}`)
}

export {HttpManager}
