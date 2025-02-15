package midware

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"Democratic_shopping_mall/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMidware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		token = strings.TrimPrefix(token, "Bearer ")

		// 解析 JWT
		username, err := utils.ParseJwt(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// 获取用户信息
		var user models.User
		if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			}
			c.Abort()
			return
		}

		// 将用户结构体存储在上下文中
		c.Set("user", user)
		c.Next()
	}
}

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前用户
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// 转换用户类型
		currentUser, ok := user.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user"})
			c.Abort()
			return
		}

		// 获取请求路径和方法
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method

		// 检查权限
		if ok, err := enforcer.Enforce(currentUser.Username, obj, act); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check permission"})
			c.Abort()
			return
		} else if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// 如果权限检查通过，继续执行
		c.Next()
	}
}
