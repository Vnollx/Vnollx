package controller

import (
	"Vnollx/common"
	"Vnollx/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()

	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(name) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
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
	if user.ID != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		log.Println("create err")
	}

	//返回结果
	ctx.JSON(500, gin.H{
		"code":    500,
		"message": "注册成功",
	})
}
