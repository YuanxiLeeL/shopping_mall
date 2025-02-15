package controllers

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"Democratic_shopping_mall/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUsername(ctx *gin.Context) {
	// 获取用户ID（从中间件中获取）
	thisuser := ctx.MustGet("user").(models.User)

	// 获取新用户名
	var requestUsername models.RequestUsername

	if err := ctx.ShouldBindJSON(&requestUsername); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//查看是否和之前的用户名相同
	if requestUsername.NewUsername == thisuser.Username {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "新用户名与旧用户名相同"})
		return
	}

	// 检查新用户名是否已存在
	var existingUser models.User
	if err := global.DB.Where("username = ?", requestUsername.NewUsername).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// 更新用户名
	thisuser.Username = requestUsername.NewUsername
	if err := global.DB.Save(&thisuser).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 返回成功消息
	ctx.JSON(http.StatusOK, gin.H{"message": "Username updated successfully"})
}

func AuthPassword(ctx *gin.Context) {

	//获得旧密码的密文
	thisuser := ctx.MustGet("user").(models.User)


	//读取用户输入的明文密码
	var oldpassword models.OldPassword
	if err := ctx.ShouldBindJSON(&oldpassword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	//验证旧密码 还有一些debug用的日志，懒得删了
	if err := bcrypt.CompareHashAndPassword([]byte(thisuser.Password), []byte(oldpassword.OldPassword)); err != nil {
		log.Printf("Hashed password from database: %s", thisuser.Password)
		log.Printf("User input old password: %s", oldpassword.OldPassword)
		debugHash, err1 := bcrypt.GenerateFromPassword([]byte(oldpassword.OldPassword), 12)
		if err1 != nil {
			ctx.JSON(111, gin.H{"error": err1.Error()})
		}
		log.Printf("hashed input password :" + string(debugHash))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	//验证通过
	ctx.JSON(http.StatusOK, gin.H{"message": "Succesfully pass the auth."})
	ctx.Next()
}

func UpdatePassword(ctx *gin.Context) {
	thisuser := ctx.MustGet("user").(models.User)

	//获取新密码
	var passwordRequest models.NewPassword
	if err := ctx.ShouldBindJSON(&passwordRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成新密码的哈希值
	newPasswordHash, err := utils.HashPassword(passwordRequest.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新密码
	thisuser.Password = newPasswordHash
	if err := global.DB.Save(&thisuser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})

}

// UpdateUserInfo 方法
func UpdateUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"username": "未登录"})
		return
	}
	//验证token
	username, err := utils.ParseJwt(token)
	if err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"token": token,
			"error": "无效的token"})
		return
	}

	// 解析请求中的 JSON 数据
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	// 更新数据库中的用户信息
	if result := global.DB.Model(&models.User{}).Where("username = ?", username).Updates(models.User{Password: hashedPassword, PhoneNum: user.PhoneNum, Email: user.Email}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	 // 清除所有 Redis 缓存
    if _, err := global.RedisDB.FlushDB().Result(); err != nil {
        c.JSON(500, gin.H{"error": "Failed to reset Redis cache"})
        return
    }

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

