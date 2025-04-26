package router

import (
	"net/http"

	"music-server-gin/global"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	if global.CONFIG.UploadConfig.SongDir != "" {
		r.StaticFS("/source/song", http.Dir(global.CONFIG.UploadConfig.SongDir))
	}

	ImgGroup := r.Group("/img")
	{
		ImgGroup.StaticFS("/singer", http.Dir(global.CONFIG.UploadConfig.SingerPicDir))
		ImgGroup.StaticFS("/songPic", http.Dir(global.CONFIG.UploadConfig.SongPicDir))
		ImgGroup.StaticFS("/songListPic", http.Dir(global.CONFIG.UploadConfig.SongListPicDir))
		ImgGroup.StaticFS("/avatar", http.Dir(global.CONFIG.UploadConfig.AvatarDir))
		ImgGroup.StaticFS("/swiper", http.Dir(global.CONFIG.UploadConfig.BannerPicDir))
	}

	// 歌曲相关路由
	SongGroup := r.Group("/song")
	{
		SongGroup.POST("/add", SongRouterApp.AddSong)
		SongGroup.DELETE("/delete", SongRouterApp.DeleteSong)
		SongGroup.GET("/all", SongRouterApp.AllSong)
		SongGroup.GET("/detail", SongRouterApp.SongOfId)
		SongGroup.GET("/singer/detail", SongRouterApp.SongOfSingerId)
		SongGroup.GET("/singerName/detail", SongRouterApp.SongOfSingerName)
		SongGroup.POST("/update", SongRouterApp.UpdateSongMsg)
		SongGroup.POST("/img/update", SongRouterApp.UpdateSongPic)
		SongGroup.POST("/url/update", SongRouterApp.UpdateSongUrl)
		SongGroup.POST("/lrc/update", SongRouterApp.UpdateSongLrc)
		SongGroup.GET("/search", SongRouterApp.SearchSongs)
	}

	// 歌手相关路由
	SingerGroup := r.Group("/singer")
	{
		SingerGroup.POST("/add", SingerRouterApp.AddSinger)
		SingerGroup.DELETE("/delete", SingerRouterApp.DeleteSinger)
		SingerGroup.GET("/all", SingerRouterApp.AllSinger)
		SingerGroup.GET("/name/detail", SingerRouterApp.SingerOfName)
		SingerGroup.GET("/sex/detail", SingerRouterApp.SingerOfSex)
		SingerGroup.POST("/update", SingerRouterApp.UpdateSingerMsg)
		SingerGroup.POST("/avatar/update", SingerRouterApp.UpdateSingerPic)
	}

	// 歌单相关路由
	SongListGroup := r.Group("/songList")
	{
		SongListGroup.GET("/all", SongListRouterApp.AllSongList)
		SongListGroup.POST("/add", SongListRouterApp.AddSongList)
		SongListGroup.GET("/delete", SongListRouterApp.DeleteSongList)
		SongListGroup.GET("/likeTitle/detail", SongListRouterApp.SongListOfTitle)
		SongListGroup.GET("/style/detail", SongListRouterApp.SongListOfStyle)
		SongListGroup.POST("/update", SongListRouterApp.UpdateSongListMsg)
		SongListGroup.POST("/img/update", SongListRouterApp.UpdateSongListPic)
	}

	// 歌单歌曲相关路由
	ListSongGroup := r.Group("/listSong")
	{
		ListSongGroup.POST("/add", ListSongRouterApp.AddListSong)
		ListSongGroup.GET("/delete", ListSongRouterApp.DeleteListSong)
		ListSongGroup.GET("/detail", ListSongRouterApp.ListSongOfSongId)
		ListSongGroup.POST("/update", ListSongRouterApp.UpdateListSongMsg)
		ListSongGroup.GET("/excle", ListSongRouterApp.GetExcle)
	}

	// 评论相关路由
	CommentGroup := r.Group("/comment")
	{
		CommentGroup.POST("/add", CommentRouterApp.AddComment)
		CommentGroup.GET("/delete", CommentRouterApp.DeleteComment)
		CommentGroup.GET("/song/detail", CommentRouterApp.CommentOfSongId)
		CommentGroup.GET("/songList/detail", CommentRouterApp.CommentOfSongListId)
		CommentGroup.POST("/like", CommentRouterApp.UpdateCommentMsg)
	}

	// 收藏相关路由
	CollectGroup := r.Group("/collection")
	{
		CollectGroup.POST("/add", CollectRouterApp.AddCollect)
		CollectGroup.DELETE("/delete", CollectRouterApp.DeleteCollect)
		CollectGroup.POST("/status", CollectRouterApp.IsCollect)
		CollectGroup.GET("/detail", CollectRouterApp.CollectOfUserId)
	}

	// 用户相关路由
	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/add", ConsumerRouterApp.AddUser)
		UserGroup.POST("/login/status", ConsumerRouterApp.LoginStatus)
		UserGroup.POST("/email/status", ConsumerRouterApp.LoginEmailStatus)
		UserGroup.POST("/resetPassword", ConsumerRouterApp.ResetPassword)
		UserGroup.GET("/sendVerificationCode", ConsumerRouterApp.SendCode)
		UserGroup.GET("/all", ConsumerRouterApp.AllUser)
		UserGroup.GET("/detail", ConsumerRouterApp.UserOfId)
		UserGroup.GET("/delete", ConsumerRouterApp.DeleteUser)
		UserGroup.POST("/update", ConsumerRouterApp.UpdateUserMsg)
		UserGroup.POST("/updatePassword", ConsumerRouterApp.UpdatePassword)
		UserGroup.POST("/avatar/update", ConsumerRouterApp.UpdateUserPic)
	}

	// 轮播图相关路由
	BannerGroup := r.Group("/banner")
	{
		BannerGroup.GET("/getAllBanner", BannerRouterApp.AllBanner)
	}

	// 文件下载相关路由
	DownloadGroup := r.Group("/download")
	{
		DownloadGroup.GET("/:fileName", FileDownloadRouterApp.DownloadFile)
	}

	// 管理员相关路由
	AdminGroup := r.Group("/admin")
	{
		AdminGroup.POST("/login/status", AdminRouterApp.LoginStatus)
	}

	return r
}
