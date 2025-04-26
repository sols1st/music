package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongRouter struct{}

func (s *SongRouter) AddSong(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	if err := service.SongServiceApp.AddSong(&song); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(song))
}

func (s *SongRouter) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	if err := service.SongServiceApp.DeleteSong(uint(id)); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("删除成功"))
}

func (s *SongRouter) AllSong(c *gin.Context) {
	songs, err := service.SongServiceApp.AllSong()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songs))
}

func (s *SongRouter) SongOfId(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	song, err := service.SongServiceApp.SongOfId(uint(id))
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(song))
}

func (s *SongRouter) SongOfSingerId(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	songs, err := service.SongServiceApp.SongOfSingerId(uint(id))
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songs))
}

func (s *SongRouter) SongOfSingerName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, BadRequest("歌手名称不能为空"))
		return
	}
	songs, err := service.SongServiceApp.SongOfSingerName(name)
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songs))
}

func (s *SongRouter) UpdateSongMsg(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	if err := service.SongServiceApp.UpdateSongMsg(&song); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("更新成功"))
}

func (s *SongRouter) UpdateSongPic(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	file, err := c.FormFile("pic")
	if err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}

	if err := service.SongServiceApp.UpdateSongPic(uint(id), fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("更新成功"))
}

func (s *SongRouter) UpdateSongUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	file, err := c.FormFile("url")
	if err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}

	if err := service.SongServiceApp.UpdateSongUrl(uint(id), fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("更新成功"))
}

func (s *SongRouter) UpdateSongLrc(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, BadRequest("无效的ID"))
		return
	}
	file, err := c.FormFile("lrc")
	if err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}

	if err := service.SongServiceApp.UpdateSongLrc(uint(id), fileBytes); err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success("更新成功"))
}

func (s *SongRouter) SearchSongs(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(400, BadRequest("搜索关键词不能为空"))
		return
	}

	songs, err := service.SongServiceApp.SearchSongs(keyword)
	if err != nil {
		c.JSON(500, Error(err.Error()))
		return
	}
	c.JSON(200, Success(songs))
}

var SongRouterApp = new(SongRouter)
