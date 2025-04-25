package router

import (
	"strconv"

	"github.com/MephistoSolsist/mysql-practice/model"
	"github.com/MephistoSolsist/mysql-practice/service"
	"github.com/gin-gonic/gin"
)

type MusicRouter struct{}

func (MusicRouter) GetMusicList(c *gin.Context) {
	var musicList []model.Music
	err := service.MusicServiceApp.GetMusicList(&musicList)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, musicList)
}

func (MusicRouter) GetMusicById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	music, err := service.MusicServiceApp.GetMusicById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, music)
}

func (MusicRouter) Update(c *gin.Context) {
	var m model.Music
	c.BindJSON(&m)
	err := service.MusicServiceApp.UpdateMusic(&m)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "succ"})
}

func (MusicRouter) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.MusicServiceApp.DeleteMusic(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "succ"})
}

func (MusicRouter) Upload(c *gin.Context) {
	var m model.Music
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = service.MusicServiceApp.UploadMusic(&m)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "succ"})
}

var MusicRouterApp = new(MusicRouter)
