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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.SongServiceApp.AddSong(&song); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, song)
}

func (s *SongRouter) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := service.SongServiceApp.DeleteSong(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "delete success"})
}

func (s *SongRouter) AllSong(c *gin.Context) {
	songs, err := service.SongServiceApp.AllSong()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songs)
}

func (s *SongRouter) SongOfId(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	song, err := service.SongServiceApp.SongOfId(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, song)
}

func (s *SongRouter) SongOfSingerId(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	songs, err := service.SongServiceApp.SongOfSingerId(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songs)
}

func (s *SongRouter) SongOfSingerName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "name is required"})
		return
	}
	songs, err := service.SongServiceApp.SongOfSingerName(name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, songs)
}

func (s *SongRouter) UpdateSongMsg(c *gin.Context) {
	var song model.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.SongServiceApp.UpdateSongMsg(&song); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

func (s *SongRouter) UpdateSongPic(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	file, err := c.FormFile("pic")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := service.SongServiceApp.UpdateSongPic(uint(id), fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

func (s *SongRouter) UpdateSongUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	file, err := c.FormFile("url")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := service.SongServiceApp.UpdateSongUrl(uint(id), fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

func (s *SongRouter) UpdateSongLrc(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	file, err := c.FormFile("lrc")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fileData, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	fileBytes := make([]byte, file.Size)
	if _, err := fileData.Read(fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := service.SongServiceApp.UpdateSongLrc(uint(id), fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

var SongRouterApp = new(SongRouter)
