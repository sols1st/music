package router

import (
	"music-server-gin/model"
	"music-server-gin/service"
	"music-server-gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	service *service.UserService
}

var UserRouterApp = &UserRouter{
	service: service.UserServiceApp,
}

// AddUser 添加用户
func (c *UserRouter) AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	if err := c.service.AddUser(&user); err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("添加成功", user))
}

// DeleteUser 删除用户
func (c *UserRouter) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	if err := c.service.DeleteUser(uint(id)); err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("删除成功", nil))
}

// AllUser 获取所有用户
func (c *UserRouter) AllUser(ctx *gin.Context) {
	users, err := c.service.AllUser()
	if err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.JSON(200, Success(users))
}

// UserOfId 根据ID获取用户
func (c *UserRouter) UserOfId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	user, err := c.service.UserOfId(uint(id))
	if err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.JSON(200, Success(user))
}

// UpdateUserMsg 更新用户信息
func (c *UserRouter) UpdateUserMsg(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	if err := c.service.UpdateUserMsg(&user); err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("更新成功", user))
}

// UpdateUserAvatar 更新用户头像
func (c *UserRouter) UpdateUserAvatar(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	if idStr == "" {
		idStr = ctx.Query("id")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	fileName := utils.GenerateUUID() + file.Filename
	if err := ctx.SaveUploadedFile(file, "upload/avatar/"+fileName); err != nil {
		ctx.JSON(400, Error("上传失败"))
		return
	}
	if err := c.service.UpdateUserAvatar(uint(id), "/img/avatar/"+fileName); err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.JSON(200, SuccessWithMessage("更新成功", nil))
}

// LoginStatus 登录状态检查
func (c *UserRouter) LoginStatus(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	user, err := c.service.LoginStatus(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.Set("userId", user.ID)
	ctx.Set("username", user.Username)
	ctx.JSON(200, SuccessWithMessage("登录成功", user))
}

// LoginEmailStatus 邮箱登录状态检查
func (c *UserRouter) LoginEmailStatus(ctx *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, Error("参数错误"))
		return
	}
	user, err := c.service.LoginEmailStatus(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(400, Error(err.Error()))
		return
	}
	ctx.Set("userId", user.ID)
	ctx.Set("username", user.Username)
	ctx.JSON(200, SuccessWithMessage("登录成功", user))
}

// UpdatePassword 更新密码
func (c *UserRouter) UpdatePassword(ctx *gin.Context) {
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
func (c *UserRouter) UpdatePassword01(ctx *gin.Context) {
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
func (c *UserRouter) ResetPassword(ctx *gin.Context) {
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
func (c *UserRouter) SendVerificationCode(ctx *gin.Context) {
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
