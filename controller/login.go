package controller

import (
	"Vnollx/common"
	"Vnollx/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(ctx *gin.Context) {

	db := common.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "密码错误",
		})
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}
