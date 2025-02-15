package controllers

import (
	"Democratic_shopping_mall/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	Likekey := "article:" + articleId + ":likes"

	if err := global.RedisDB.Incr(Likekey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully like the article"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleId := ctx.Param("id")

	Likekey := "article:" + articleId + ":likes"

	Likes, err := global.RedisDB.Get(Likekey).Result()

	if err == redis.Nil {
		Likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": Likes})

}
