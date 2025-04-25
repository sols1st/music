package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SingerRouter struct{}

func (s *SingerRouter) AddSinger(c *gin.Context) {
	var singer model.Singer
	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.SingerServiceApp.AddSinger(&singer); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, singer)
}

func (s *SingerRouter) DeleteSinger(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := service.SingerServiceApp.DeleteSinger(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "delete success"})
}

func (s *SingerRouter) UpdateSingerMsg(c *gin.Context) {
	var singer model.Singer
	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.SingerServiceApp.UpdateSingerMsg(&singer); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

func (s *SingerRouter) UpdateSingerPic(c *gin.Context) {
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

	if err := service.SingerServiceApp.UpdateSingerPic(uint(id), fileBytes); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

func (s *SingerRouter) AllSinger(c *gin.Context) {
	singers, err := service.SingerServiceApp.AllSinger()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": singers})
}

func (s *SingerRouter) SingerOfName(c *gin.Context) {
	name := c.Query("name")
	singers, err := service.SingerServiceApp.SingerOfName(name)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": singers})
}

func (s *SingerRouter) SingerOfSex(c *gin.Context) {
	sex := c.Query("sex")
	singers, err := service.SingerServiceApp.SingerOfSex(sex)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "获取失败"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": singers})
}

var SingerRouterApp = new(SingerRouter)
