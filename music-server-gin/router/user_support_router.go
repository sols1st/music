package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserSupportRouter struct{}

func (s *UserSupportRouter) AddUserSupport(c *gin.Context) {
	var userSupport model.UserSupport
	if err := c.ShouldBindJSON(&userSupport); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.UserSupportServiceApp.AddUserSupport(&userSupport); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userSupport)
}

func (s *UserSupportRouter) DeleteUserSupport(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := service.UserSupportServiceApp.DeleteUserSupport(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "delete success"})
}

func (s *UserSupportRouter) AllUserSupport(c *gin.Context) {
	userSupports, err := service.UserSupportServiceApp.AllUserSupport()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userSupports)
}

func (s *UserSupportRouter) UserSupportOfUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid userId"})
		return
	}
	userSupports, err := service.UserSupportServiceApp.UserSupportOfUserId(uint(userId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userSupports)
}

func (s *UserSupportRouter) UpdateUserSupportMsg(c *gin.Context) {
	var userSupport model.UserSupport
	if err := c.ShouldBindJSON(&userSupport); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.UserSupportServiceApp.UpdateUserSupportMsg(&userSupport); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "update success"})
}

var UserSupportRouterApp = new(UserSupportRouter)
