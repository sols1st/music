package router

import (
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

// AddUser 添加用户
func (c *ConsumerRouter) AddUser(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.AddUser(&consumer); err != nil {
		ctx.JSON(500, Error("注册失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("注册成功", nil))
}

// DeleteUser 删除用户
func (c *ConsumerRouter) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.DeleteUser(uint(id)); err != nil {
		ctx.JSON(500, Error("删除失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("删除成功", nil))
}

// AllUser 获取所有用户
func (c *ConsumerRouter) AllUser(ctx *gin.Context) {
	consumers, err := c.service.AllUser()
	if err != nil {
		ctx.JSON(500, Error("获取失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success(consumers))
}

// UserOfId 根据ID获取用户
func (c *ConsumerRouter) UserOfId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	consumer, err := c.service.UserOfId(uint(id))
	if err != nil {
		ctx.JSON(500, Error("获取失败: "+err.Error()))
		return
	}
	ctx.JSON(200, Success(consumer))
}

// UpdateUserMsg 更新用户信息
func (c *ConsumerRouter) UpdateUserMsg(ctx *gin.Context) {
	var consumer model.Consumer
	if err := ctx.ShouldBindJSON(&consumer); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.UpdateUserMsg(&consumer); err != nil {
		ctx.JSON(500, Error("更新失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("更新成功", nil))
}

// UpdateUserAvatar 更新用户头像
func (c *ConsumerRouter) UpdateUserAvatar(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, BadRequest("文件上传失败: "+err.Error()))
		return
	}

	// 处理文件上传
	avatarPath := "img/avatorImages/" + file.Filename
	if err := ctx.SaveUploadedFile(file, avatarPath); err != nil {
		ctx.JSON(500, Error("文件保存失败: "+err.Error()))
		return
	}

	if err := c.service.UpdateUserAvatar(uint(id), avatarPath); err != nil {
		ctx.JSON(500, Error("更新失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("更新成功", nil))
}

// LoginStatus 登录状态检查
func (c *ConsumerRouter) LoginStatus(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	consumer, err := c.service.LoginStatus(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(401, Unauthorized(err.Error()))
		return
	}

	// 设置session
	ctx.Set("userId", consumer.ID)
	ctx.Set("username", consumer.Username)

	ctx.JSON(200, SuccessWithMessage("登录成功", consumer))
}

// LoginEmailStatus 邮箱登录状态检查
func (c *ConsumerRouter) LoginEmailStatus(ctx *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	consumer, err := c.service.LoginEmailStatus(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(401, Unauthorized(err.Error()))
		return
	}

	// 设置session
	ctx.Set("userId", consumer.ID)
	ctx.Set("username", consumer.Username)

	ctx.JSON(200, SuccessWithMessage("登录成功", consumer))
}

// UpdatePassword 更新密码
func (c *ConsumerRouter) UpdatePassword(ctx *gin.Context) {
	var updateRequest struct {
		ID          uint   `json:"id"`
		OldPassword string `json:"oldPassword"`
		Password    string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.UpdatePassword(updateRequest.ID, updateRequest.OldPassword, updateRequest.Password); err != nil {
		ctx.JSON(500, Error("更新失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("更新成功", nil))
}

// UpdatePassword01 更新密码（简化版）
func (c *ConsumerRouter) UpdatePassword01(ctx *gin.Context) {
	var updateRequest struct {
		ID       uint   `json:"id"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.UpdatePassword01(updateRequest.ID, updateRequest.Password); err != nil {
		ctx.JSON(500, Error("更新失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("更新成功", nil))
}

// ResetPassword 重置密码
func (c *ConsumerRouter) ResetPassword(ctx *gin.Context) {
	var resetRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&resetRequest); err != nil {
		ctx.JSON(400, BadRequest("参数错误: "+err.Error()))
		return
	}

	if err := c.service.ResetPassword(resetRequest.Email, resetRequest.Password); err != nil {
		ctx.JSON(500, Error("重置失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("重置成功", nil))
}

// SendVerificationCode 发送验证码
func (c *ConsumerRouter) SendVerificationCode(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(400, BadRequest("邮箱不能为空"))
		return
	}

	if err := c.service.SendVerificationCode(email); err != nil {
		ctx.JSON(500, Error("发送失败: "+err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("发送成功", nil))
}
