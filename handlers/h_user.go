/*
   @TIME: 2023/6/13 12:20
   @Author: Nightz
   @Site:
   @File: h_user.go
   @SoftWare: GoLand
*/

package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"yishi-ai.com/go-sm/constant"
	model "yishi-ai.com/go-sm/models"
	"yishi-ai.com/go-sm/utils"
)

type userLogin struct {
	UserName string `json:"username" form:"username" validate:"required"`
	PassWord string `json:"password" form:"password" validate:"required"`
}

type userRegister struct {
	UserName string `json:"username" form:"username" validate:"required"`
	PassWord string `json:"password" form:"password" validate:"required"`
}

// UserLoginHandler 用户登录
func UserLoginHandler(c *gin.Context) {
	login := userLogin{}
	user := model.User{}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数校验失败"})
		return
	}
	if userNotFound := constant.DB.Model(&model.User{}).Where(login).First(&user).Error; userNotFound == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"message": "用户信息不存在"})
	} else {
		jwt, _ := utils.GenerateJWT(strconv.Itoa(int(user.ID)))
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"token":   jwt,
		})
	}

}

// UserInfoHandler 用户信息
func UserInfoHandler(c *gin.Context) {
	user := &model.User{}
	userID, _ := c.Get("userID")
	if err := constant.DB.Model(&model.User{}).Where("id = ?", userID).First(user).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{"message": "该用户已登录"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "未登录"})
	}
}

// UserRegisterHandler 用户注册
func UserRegisterHandler(c *gin.Context) {
	user := &model.User{}
	register := userRegister{}
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数校验失败"})
		return
	}
	userQuery := constant.DB.Model(&model.User{})

	if userNotFound := userQuery.First(&user).Error; userNotFound == gorm.ErrRecordNotFound {
		constant.DB.Model(&model.User{}).Create(&model.User{
			UserName: register.UserName,
			PassWord: register.PassWord,
		})
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
			"data":    register,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "已存在同名用户"})
	}
}

/*

 */
