package router

import (
	"time"

	"github.com/MephistoSolsist/mysql-practice/model"
	"github.com/MephistoSolsist/mysql-practice/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (UserRouter) Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	u.GmtCreate = time.Now().Format("2006-01-02 15:04:05")
	u.Role = "user"
	err = service.UserServiceApp.Register(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, u)

}

func (UserRouter) Login(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	u, err = service.UserServiceApp.Login(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, u)

}

func (UserRouter) Delete(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = service.UserServiceApp.Delete(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "delete success"})

}

var UserRouterApp = new(UserRouter)
