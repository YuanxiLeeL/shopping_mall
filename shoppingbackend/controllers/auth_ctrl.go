package controllers

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"Democratic_shopping_mall/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	token, err := utils.GenarateJwt(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	if err := global.DB.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tables"})
		return
	}

	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := global.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("Invalid password")})
		return
	}

	token, err := utils.GenarateJwt(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":"10000",
		"statusText":"登录成功",
		"token": token,
	},
	)
}



// GetCurrentUserInfo 获取当前用户信息
func GetCurrentUserInfo(c *gin.Context) {
	
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"username": "未登录"})
		return
	}

	//验证token
	username, err := utils.ParseJwt(token)
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"error": "无效的token"})
		return
	}

	// 根据用户名查询用户信息
    var user models.User
    if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusOK, gin.H{"error": "用户未找到"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "查询用户信息失败"})
        }
        return
    }

	// 去除密码字段
    userWithoutPassword := models.UserWithoutPassword{
        Username: user.Username,
        Email:    user.Email,
        PhoneNum: user.PhoneNum,
    }


	c.JSON(http.StatusOK, gin.H{
		"data" : userWithoutPassword,
		"status":"10000",
		"isLoggedIn": true,})

	
}
