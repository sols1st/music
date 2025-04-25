package router

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"music-server-gin/global"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 修改CORS配置，使用更宽松的设置
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 添加静态文件服务 - 简化成一个总的文件服务
	uploadDir, err := filepath.Abs(global.CONFIG.UploadConfig.UploadDir)
	if err != nil {
		log.Printf("获取上传目录绝对路径失败: %v", err)
		uploadDir = "./upload" // 使用默认路径
	}

	// 日志输出调试信息
	log.Printf("静态文件目录: %s", uploadDir)
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		log.Printf("警告: 上传目录不存在: %s", uploadDir)
	}

	// // 提供静态文件服务
	// r.StaticFS("/img", http.Dir(uploadDir))

	// 配置各种图片目录的静态文件服务
	if global.CONFIG.UploadConfig.SingerPicDir != "" {
		r.StaticFS("/img/singer", http.Dir(global.CONFIG.UploadConfig.SingerPicDir))
	}
	if global.CONFIG.UploadConfig.SongPicDir != "" {
		r.StaticFS("/img/songPic", http.Dir(global.CONFIG.UploadConfig.SongPicDir))
	}
	if global.CONFIG.UploadConfig.SongListPicDir != "" {
		r.StaticFS("/img/songListPic", http.Dir(global.CONFIG.UploadConfig.SongListPicDir))
	}
	if global.CONFIG.UploadConfig.AvatarDir != "" {
		r.StaticFS("/img/avatar", http.Dir(global.CONFIG.UploadConfig.AvatarDir))
	}
	if global.CONFIG.UploadConfig.BannerPicDir != "" {
		r.StaticFS("/img/swiper", http.Dir(global.CONFIG.UploadConfig.BannerPicDir))
	}

	// 歌曲相关路由
	SongGroup := r.Group("/song")
	{
		SongGroup.POST("/add", SongRouterApp.AddSong)
		SongGroup.DELETE("/delete", SongRouterApp.DeleteSong)
		SongGroup.GET("/", SongRouterApp.AllSong)
		SongGroup.GET("/detail", SongRouterApp.SongOfId)
		SongGroup.GET("/singer/detail", SongRouterApp.SongOfSingerId)
		SongGroup.GET("/singerName/detail", SongRouterApp.SongOfSingerName)
		SongGroup.POST("/update", SongRouterApp.UpdateSongMsg)
		SongGroup.POST("/img/update", SongRouterApp.UpdateSongPic)
		SongGroup.POST("/url/update", SongRouterApp.UpdateSongUrl)
		SongGroup.POST("/lrc/update", SongRouterApp.UpdateSongLrc)
	}

	// 歌手相关路由
	SingerGroup := r.Group("/singer")
	{
		SingerGroup.POST("/add", SingerRouterApp.AddSinger)
		SingerGroup.DELETE("/delete", SingerRouterApp.DeleteSinger)
		SingerGroup.GET("/", SingerRouterApp.AllSinger)
		SingerGroup.GET("/name/detail", SingerRouterApp.SingerOfName)
		SingerGroup.GET("/sex/detail", SingerRouterApp.SingerOfSex)
		SingerGroup.POST("/update", SingerRouterApp.UpdateSingerMsg)
		SingerGroup.POST("/avatar/update", SingerRouterApp.UpdateSingerPic)
	}

	// 歌单相关路由
	SongListGroup := r.Group("/songList")
	{
		SongListGroup.POST("/add", SongListRouterApp.AddSongList)
		SongListGroup.GET("/delete", SongListRouterApp.DeleteSongList)
		SongListGroup.GET("/", SongListRouterApp.AllSongList)
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

	// 评分相关路由
	RankListGroup := r.Group("/rankList")
	{
		RankListGroup.POST("/add", RankListRouterApp.AddRank)
		RankListGroup.GET("/", RankListRouterApp.RankOfSongListId)
		RankListGroup.GET("/user", RankListRouterApp.RankOfConsumerId)
	}

	// 用户相关路由
	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/add", ConsumerRouterApp.AddUser)
		UserGroup.POST("/login/status", ConsumerRouterApp.LoginStatus)
		UserGroup.POST("/email/status", ConsumerRouterApp.LoginEmailStatus)
		UserGroup.POST("/resetPassword", ConsumerRouterApp.ResetPassword)
		UserGroup.GET("/sendVerificationCode", ConsumerRouterApp.SendCode)
		UserGroup.GET("/", ConsumerRouterApp.AllUser)
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

	// 用户支持相关路由
	UserSupportGroup := r.Group("/userSupport")
	{
		UserSupportGroup.POST("/add", UserSupportRouterApp.AddUserSupport)
		UserSupportGroup.DELETE("/delete", UserSupportRouterApp.DeleteUserSupport)
		UserSupportGroup.GET("/", UserSupportRouterApp.AllUserSupport)
		UserSupportGroup.GET("/user/detail", UserSupportRouterApp.UserSupportOfUserId)
		UserSupportGroup.POST("/update", UserSupportRouterApp.UpdateUserSupportMsg)
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
