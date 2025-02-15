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

func AddToCart(c *gin.Context) {
	// 获取用户ID（从中间件中获取）
	// userName := c.MustGet("username").(string)
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

	
	
	var cartItem models.CartItem

	if err := global.DB.AutoMigrate(&cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	goodName := c.Param("goodname")

	// 获取商品ID和数量
	// var request struct {
	// 	Quantity *int  `json:"quantity"`
	// }
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if request.Quantity != nil{
	// 	quantity = *request.Quantity
	// }

	// 检查商品是否存在
	var good models.Good
	if err := global.DB.Where("name = ?", goodName).First(&good).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Good not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get good"})
		}
		return
	}

	//检索价格
	if err := global.DB.Where("name =?", goodName).First(&good).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get good"})
	}
	theGood := models.Good{
		Price: good.Price,
	}

	// 检查购物车中是否已存在该商品
	if err := global.DB.Where("username = ? AND good_name = ?", username, goodName).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果购物车中没有该商品，创建一个新的购物车项
			cartItem = models.CartItem{
				Username:   username,
				GoodName:   goodName,
				Quantity: 1,
				Price: theGood.Price,
			}
			if err := global.DB.Create(&cartItem).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
				return
			}
		} else{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error"})}
	}else{
		// 如果购物车中已存在该商品，更新数量
			cartItem.Quantity += 1
			if err := global.DB.Save(&cartItem).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
				return
			}


	}
	// 返回添加到购物车的商品项
	c.JSON(http.StatusOK, cartItem)
}

func RemoveFromCart(c *gin.Context) {
	// 获取用户ID（从中间件中获取）
	// userName := c.MustGet("username").(string)

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
	// 获取商品ID
	goodName := c.Param("goodname")

	// 删除购物车中的商品项
	if err := global.DB.Where("username = ? AND good_name = ?", username, goodName).Unscoped().Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove from cart"})
		return
	}

	 // 清除所有 Redis 缓存
    if _, err := global.RedisDB.FlushDB().Result(); err != nil {
        c.JSON(500, gin.H{"error": "Failed to reset Redis cache"})
        return
    }

	// 返回成功消息
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}

func GetCart(c *gin.Context) {
	// 获取用户名（从中间件中获取）
	// userName := c.MustGet("username").(string)
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
	// 获取购物车中的商品列表
	var cartItems []models.CartItem
	if err := global.DB.Where("username = ?", username).Find(&cartItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart"})
		}
		return
	}

	// 返回购物车中的商品列表
	c.JSON(http.StatusOK, cartItems)
}

func PlaceOrder(c *gin.Context) {
	// 获取用户名（从中间件中获取）
	// userName := c.MustGet("username").(string)
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

	var cartItems []models.CartItem
    // 根据用户名查找购物车商品
    if err := global.DB.Where("username = ?", username).Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "查询购物车信息失败"})
        return
    }

	 totalAmount := int64(0)
    // 计算购物车内所有商品的价格总和，每个商品价格乘以其数量
    for _, item := range cartItems {
        totalAmount += item.Price * int64(item.Quantity)
    }

	// 清空购物车
    if err := global.DB.Where("username = ?", username).Unscoped().Delete(&models.CartItem{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "清空购物车失败"})
        return
    }

	// 清除所有 Redis 缓存
    if _, err := global.RedisDB.FlushDB().Result(); err != nil {
        c.JSON(500, gin.H{"error": "Failed to reset Redis cache"})
        return
    }

	 // 构造返回数据
    response := gin.H{
        "totalAmount": totalAmount,
    }
    c.JSON(http.StatusOK, response)



}