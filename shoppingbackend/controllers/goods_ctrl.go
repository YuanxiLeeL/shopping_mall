package controllers

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var cacheKeyforG = "goods"

func CreateGood(ctx *gin.Context) {
	var good models.Good

	if err := ctx.ShouldBindJSON(&good); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.DB.AutoMigrate(&good); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tables"})
		return
	}

	if err := global.DB.Create(&good).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	if err := global.RedisDB.Del(cacheKeyforG).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, good)

}
//ai修改前内容
// func GetGoods(ctx *gin.Context) {
// 	var good []models.Good

// 	cacheData, err := global.RedisDB.Get(cacheKeyforG).Result()

// 	if err == redis.Nil {
// 		if err := global.DB.Find(&good).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			} else {
// 				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
// 			}
// 			return
// 		}
// 		GoodJson, err := json.Marshal(good)
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
// 			return
// 		}

// 		if err := global.RedisDB.Set(cacheKeyforG, GoodJson, 5*time.Minute).Err(); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, good)

// 	} else if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
// 		return
// 	} else {
// 		var good []models.Good
// 		if err := json.Unmarshal([]byte(cacheData), &good); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// goods := models.Goods{
// 		// 	Goods: good,
// 		// 	Total: int64(len(good)),
// 		// }
// 		ctx.JSON(http.StatusOK, good)
// 	}

// }

//ai修改后内容
func GetGoods(ctx *gin.Context) {
	var good []models.Good

	// 获取查询参数 goodName
	goodName := ctx.Query("goodName")

	// 构造缓存键
	cacheKey := "goods"
	if goodName != "" {
		cacheKey = "goods:" + goodName
	}

	// 尝试从 Redis 缓存中获取数据
	cacheData, err := global.RedisDB.Get(cacheKey).Result()

	if err == redis.Nil {
		// 如果缓存中没有数据，从数据库中查询
		if goodName != "" {
			// 如果传入了 goodName 参数，根据商品名称查询
			if err := global.DB.Where("name = ?", goodName).Find(&good).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					ctx.JSON(http.StatusNotFound, gin.H{"error": "Good not found"})
					return
				}
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get goods"})
				return
			}
		} else {
			// 如果没有传入 goodName 参数，查询所有商品
			if err := global.DB.Find(&good).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					ctx.JSON(http.StatusNotFound, gin.H{"error": "Goods not found"})
					return
				}
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get goods"})
				return
			}
		}

		// 将查询结果序列化为 JSON
		GoodJson, err := json.Marshal(good)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 将数据存入 Redis 缓存
		if err := global.RedisDB.Set(cacheKey, GoodJson, 5*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 返回查询结果
		ctx.JSON(http.StatusOK, good)
	} else if err != nil {
		// 如果从缓存中获取数据时出错
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		// 如果缓存中有数据，直接解析并返回
		if err := json.Unmarshal([]byte(cacheData), &good); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, good)
	}
}

// func GetGoodsByName(ctx *gin.Context) {
// 	name := ctx.Param("name")
// 	var good models.Good

// 	if err := global.DB.Where("name = ?", name).First(&good).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		} else {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get good"})
// 		}

// 	}

// }

func DelGoodsbyName(ctx *gin.Context) {
	name := ctx.Param("name")

	if err := global.DB.Unscoped().Where("name = ?", name).Delete(&models.Good{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	 // 清除所有 Redis 缓存
    if _, err := global.RedisDB.FlushDB().Result(); err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to reset Redis cache"})
        return
    }
	ctx.JSON(http.StatusOK, gin.H{"message": "Good deleted successfully"})

}

func UpdateGood(c *gin.Context) {
	// 获取商品 ID
	goodID := c.Param("id")

	// 获取请求中的商品信息
	var updateData models.Good
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询商品
	var good models.Good
	if err := global.DB.Where("id = ?", goodID).First(&good).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Good not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get good"})
		}
		return
	}

	// 更新商品信息
	if err := global.DB.Model(&models.Good{}).Where("id = ?", goodID).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update good"})
		return
	}

	// 返回更新后的商品信息
	c.JSON(http.StatusOK, gin.H{"message": "Good updated successfully"})
}

func SearchGoods(ctx *gin.Context) {
	// 获取搜索关键词
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	// 查询商品
	var goods []models.Good
	if err := global.DB.Where("name LIKE ?", "%"+query+"%").Find(&goods).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search goods"})
		return
	}

	// 返回查询结果
	ctx.JSON(http.StatusOK, goods)
}

// GetSingleGoodInfo 获取单个商品信息及其评论
func GetSingleGoodInfo(c *gin.Context) {
    // token := c.GetHeader("Authorization")
	// if token == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"username": "未登录"})
	// 	return
	// }
	
	// 获取商品名称
    goodName := c.Param("name")
	//验证token
	// username, err := utils.ParseJwt(token)
	// if err!= nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"token": token,
	// 		"error": "无效的token"})
	// 	return
	// }
    // 获取商品信息
    var good models.Good
    if err := global.DB.Where("name = ?", goodName).First(&good).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Good not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get good"})
        return
    }

    // 获取该商品下的所有评论
    var comments []models.Comment
    if err := global.DB.Where("good_name = ?", goodName).Find(&comments).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            comments = []models.Comment{}
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
            return
        }
    }

	// 格式化评论的创建时间
    var formattedComments []map[string]interface{}
    for _, comment := range comments {
        formattedComment := map[string]interface{}{
            "ID":        comment.ID,
            "username":  comment.UserName,
            "content":   comment.Content,
            "CreatedAt": comment.CreatedAt.Format("2006-01-02 15:04:05"),
            "goodname":  comment.GoodName,
        }
        formattedComments = append(formattedComments, formattedComment)
    }

    // 构造返回格式
    response := gin.H{
        "info": good,
        "comments": formattedComments,
    }

    c.JSON(http.StatusOK, response)
}