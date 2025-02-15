package router

import (
	"Democratic_shopping_mall/controllers"
	"Democratic_shopping_mall/midware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(midware.Cors())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "token"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))


	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.GET("/current", controllers.GetCurrentUserInfo)
		auth.POST("/register", controllers.Register)
		
	}

	api := r.Group("/api")
	api.Use(midware.AuthMidware())//midware.CasbinMiddleware(global.Enforcer)
	{

		
		//注册商品相关路由
		api.GET("/goods", controllers.GetGoods)
		api.POST("/goods", controllers.CreateGood)
		// api.GET("/goods/:name", controllers.GetGoodsByName)
		api.DELETE("/goods/:name", controllers.DelGoodsbyName)
		api.GET("/singlegood/:name", controllers.GetSingleGoodInfo)
		// api.PUT("/goods/:id", controllers.UpdateGood)

		// api.POST("/goods/:id", controllers.LikeArticle)
		// api.GET("/goods/:id", controllers.GetArticleLikes)

		// 注册评论路由
		api.POST("/singlegood/:goodname/comments", controllers.CreateComment)
		// api.GET("/goods/:good_id/comments", controllers.GetCommentsByGoodID)
		api.GET("/comments/:id", controllers.GetCommentsByGoodID)
		api.PUT("/comments/:id", controllers.UpdateComment)
		api.DELETE("/comments/:id", controllers.DeleteComment)
		api.GET("/comments/:id/permission",controllers.CheckCommentPermission)

		// 注册用户修改路由
		api.PUT("/user/info/edit", controllers.UpdateUserInfo)
		api.POST("/user/authpassword", controllers.AuthPassword)
		api.PUT("/user/password", controllers.UpdatePassword)

		// 注册分类路由
		
		api.POST("/category/:category", controllers.CreateCategory)
		api.GET("/category", controllers.GetCategories)
		api.GET("/category/:categoryname", controllers.GetGoodsByCategory)
		api.DELETE("/:id", controllers.DeleteCategory)

		api.GET("/cart", controllers.GetCart)
		api.DELETE("/cart/:goodname", controllers.RemoveFromCart)
		api.POST("/cart/:goodname", controllers.AddToCart)
		api.GET("/placeorder", controllers.PlaceOrder)
		// cart := r.Group("/api/goods/cart")
		// {
		// 	cart.POST("/", controllers.AddToCart)
		// 	cart.DELETE("/:goodname", controllers.RemoveFromCart)
		// 	cart.GET("/", controllers.GetCart)
		// }

	return r
}
}