package router

import (
	"fmt"
	"music-server-gin/model"
	"music-server-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConsumerRouter struct {
	service *service.ConsumerService
}

var ConsumerRouterApp = &ConsumerRouter{
	service: service.ConsumerServiceApp,
}

func (c *ConsumerRouter) AddUser(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	fmt.Println(consumer)
	if err := c.service.AddConsumer(&consumer); err != nil {
		ctx.JSON(500, Error(err.Error()))
		return
	}
	ctx.JSON(200, Success("添加成功"))
}

func (c *ConsumerRouter) LoginStatus(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	_, err := c.service.VerifyPassword(consumer.Username, consumer.Password)
	if err != nil {
		ctx.JSON(500, Error("用户名或密码错误"))
		return
	}
	ctx.JSON(200, Success("登录成功"))
}

func (c *ConsumerRouter) LoginEmailStatus(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	_, err := c.service.VerifyEmailPassword(consumer.Email, consumer.Password)
	if err != nil {
		ctx.JSON(500, Error("邮箱或密码错误"))
		return
	}
	ctx.JSON(200, Success("登录成功"))
}

func (c *ConsumerRouter) ResetPassword(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.ResetPassword(consumer.Email, consumer.Password); err != nil {
		ctx.JSON(500, Error("重置密码失败"))
		return
	}
	ctx.JSON(200, Success("重置密码成功"))
}

func (c *ConsumerRouter) SendCode(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.SendVerificationCode(email); err != nil {
		ctx.JSON(500, Error("发送验证码失败"))
		return
	}
	ctx.JSON(200, Success("发送验证码成功"))
}

func (c *ConsumerRouter) AllUser(ctx *gin.Context) {
	consumers, err := c.service.AllConsumer()
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	ctx.JSON(200, Success(consumers))
}

func (c *ConsumerRouter) UserOfId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	consumer, err := c.service.ConsumerOfId(uint(id))
	if err != nil {
		ctx.JSON(500, Error("获取失败"))
		return
	}
	ctx.JSON(200, Success(consumer))
}

func (c *ConsumerRouter) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.DeleteConsumer(uint(id)); err != nil {
		ctx.JSON(500, Error("删除失败"))
		return
	}
	ctx.JSON(200, Success("删除成功"))
}

func (c *ConsumerRouter) UpdateUserMsg(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.UpdateConsumerMsg(&consumer); err != nil {
		ctx.JSON(500, Error("更新失败"))
		return
	}
	ctx.JSON(200, Success("更新成功"))
}

func (c *ConsumerRouter) UpdatePassword(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.UpdatePassword(&consumer); err != nil {
		ctx.JSON(500, Error("更新密码失败"))
		return
	}
	ctx.JSON(200, Success("更新密码成功"))
}

func (c *ConsumerRouter) UpdateUserPic(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误"))
		return
	}
	if err := c.service.UpdateConsumerAvatar(uint(id), file); err != nil {
		ctx.JSON(500, Error("更新头像失败"))
		return
	}
	ctx.JSON(200, Success("更新头像成功"))
}
