package controllers

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"Democratic_shopping_mall/utils"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var cacheKeyforC = "comment"

func CreateComment(ctx *gin.Context) {
	var comment models.Comment
	goodname := ctx.Param("goodname")
token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"username": "未登录"})
		return
	}

	//验证token
	username, err := utils.ParseJwt(token)
	if err!= nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"token": token,
			"error": "无效的token"})
		return
	}
	// 绑定 JSON 数据到 comment 结构体
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserName = username
	comment.GoodName = goodname
	//验证表格是否存在
	if err := global.DB.AutoMigrate(&comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 创建评论
	if err := global.DB.Create(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// 删除 Redis 缓存（如果有的话）
	cacheKey := "comments:" + ctx.Param("good_id")
	if err := global.RedisDB.Del(cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cache"})
		return
	}

	// 返回创建的评论
	ctx.JSON(http.StatusCreated, comment)
}

func GetCommentsByGoodID(ctx *gin.Context) {
	goodID := ctx.Param("good_id")

	// 尝试从 Redis 缓存中获取评论列表
	cacheData, err := global.RedisDB.Get(cacheKeyforC).Result()
	if err == redis.Nil {
		// 缓存中没有数据，从数据库中查询
		var comments []models.Comment
		if err := global.DB.Where("good_id = ?", goodID).Find(&comments).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "No comments found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments from database"})
			}
			return
		}

		// 将评论列表序列化为 JSON
		commentsJSON, err := json.Marshal(comments)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal comments"})
			return
		}

		// 更新 Redis 缓存
		if err := global.RedisDB.Set(cacheKeyforG, commentsJSON, 5*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
			return
		}

		// 返回评论列表
		ctx.JSON(http.StatusOK, comments)
	} else if err != nil {
		// Redis 查询出错
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cache"})
		return
	} else {
		// 从缓存中获取评论列表
		var comments []models.Comment
		if err := json.Unmarshal([]byte(cacheData), &comments); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cache data"})
			return
		}

		// 返回评论列表
		ctx.JSON(http.StatusOK, comments)
	}
}

func UpdateComment(c *gin.Context) {
	// 获取评论 ID
	commentID := c.Param("id")
	// token := c.GetHeader("Authorization")
	// if token == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"username": "未登录"})
	// 	return
	// }

	// //验证token
	// username, err := utils.ParseJwt(token)
	// if err!= nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"token": token,
	// 		"error": "无效的token"})
	// 	return
	// }

	// 获取用户ID（从中间件中获取）
	
	type newcomment struct {
		Content string `json:"content"`
	}
	// 获取评论内容
	var rescomment newcomment
	if err := c.ShouldBindJSON(&rescomment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查评论是否存在
	var existingComment models.Comment
	if err := global.DB.Where("id = ?", commentID).First(&existingComment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment"})
		}
		return
	}

	// 更新评论内容和评分
	// existingComment.Content = comment.Content
	// if existingComment.ReplyTo == 0 {
	// 	// 只有针对商品的评论才允许设置评分
	// 	existingComment.Rating = comment.Rating
	// }

	// 更新评论内容
	existingComment.Content = rescomment.Content
	if err := global.DB.Save(&existingComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新评论失败"})
		return
	}

	if err := global.DB.Save(&existingComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	// 更新 Redis 缓存
	cacheKeyforC := "comments:" + c.Param("good_id")
	if err := global.RedisDB.Del(cacheKeyforC).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cache"})
		return
	}

	// 返回更新后的评论
	c.JSON(http.StatusOK, existingComment)
}

func DeleteComment(c *gin.Context) {
	// 获取评论 ID
	commentID := c.Param("id")
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"username": "未登录"})
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
	

	// 检查评论是否存在
	var comment models.Comment
	if err := global.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment"})
		}
		return
	}

	// 验证用户是否有权限删除评论
	if comment.UserName != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this comment"})
		return
	}

	// 删除评论
	if err := global.DB.Unscoped().Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	// 删除 Redis 缓存
	cacheKeyforC := "comments:" + c.Param("goodname")
	if err := global.RedisDB.Del(cacheKeyforC).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cache"})
		return
	}

	// 返回成功消息
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// CheckCommentPermission 检查用户是否有权限修改评论
func CheckCommentPermission(c *gin.Context) {
	// 获取评论 ID
	commentID := c.Param("id")

	// 获取 Token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// 验证 Token
	username, err := utils.ParseJwt(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
		return
	}

	// 获取评论信息
	var comment models.Comment
	if err := global.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "评论未找到"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}

	// 比对用户名
	if comment.UserName != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限修改这条评论"})
		return
	}

	// 如果验证通过，返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "验证通过"})
}
