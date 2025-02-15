package controllers

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建分类（仅限管理员账号）
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.DB.AutoMigrate(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	if err := global.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// 获取所有分类
func GetCategories(c *gin.Context) {
	
	var categories []models.Category
	if err := global.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func GetGoodsByCategory(c *gin.Context) {
	// 获取分类的 CategoryManualID
	categoryname := c.Param("categoryname")

	// 查询该分类下的所有商品
	var goods []models.Good
	if err := global.DB.Where("category = ?", categoryname).Find(&goods).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No goods found in this category"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get goods"})
		}
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, goods)
}

// 删除分类（仅限管理员账号）
func DeleteCategory(c *gin.Context) {
	// 获取手动设定的分类 ID
	categoryManualID := c.Param("id")

	// 查询分类
	var category models.Category
	if err := global.DB.Where("category_manual_id = ?", categoryManualID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get category"})
		}
		return
	}

	// 硬删除分类
	if err := global.DB.Unscoped().Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	// 返回删除成功的消息
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
