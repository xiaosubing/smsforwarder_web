package service

import (
	"bg/middleware"
	"bg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	recInfos := new(LoginRequest)
	err := c.ShouldBind(recInfos)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  400,
		})
		fmt.Println(err)
		return
	}

	fmt.Println("=================1111111111")
	user, err := models.GetUserInfo(recInfos.Username, recInfos.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}
	fmt.Println("==========22222222222222")
	// 生成token
	token, err := middleware.GenerateToken(user.ID, user.Username, middleware.TokenExpireTime)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}
	fmt.Println("================44444444444")
	refreshToken, err := middleware.GenerateToken(user.ID, user.Username, middleware.RefreshTokenExpireTime)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  "404",
		})
		return
	}

	date := &LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": *date,
	})
}
