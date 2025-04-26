package router

import (
	"music-server-gin/model"
	"music-server-gin/service"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (s *AdminRouter) LoginStatus(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(400, BadRequest(err.Error()))
		return
	}
	if !service.AdminServiceApp.VerifyPassword(admin.Name, admin.Password) {
		c.JSON(500, Error("用户名或密码错误"))
		return
	}
	c.JSON(200, Success("登录成功"))
}

var AdminRouterApp = new(AdminRouter)
