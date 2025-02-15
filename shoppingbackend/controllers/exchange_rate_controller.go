package controllers

import (
	"Democratic_shopping_mall/global"
	"Democratic_shopping_mall/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate

	if err := ctx.ShouldBindJSON(&exchangeRate);err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exchangeRate.Date = time.Now()

	if err := global.DB.AutoMigrate(&exchangeRate); err!= nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	if err := global.DB.Create(&exchangeRate).Error; err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Exchange rate created successfully"})

}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRate []models.ExchangeRate
	if err := global.DB.Find(&exchangeRate).Error; err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRate)
}
