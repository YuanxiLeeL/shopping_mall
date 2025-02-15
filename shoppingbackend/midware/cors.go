package midware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               // 请求方法
		origin := c.Request.Header.Get("Origin") // 请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                                                                                                                                                  // 允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                                                                                                           // 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Headers", "*")                                                                                                                                                                 // 允许的头类型
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, XMLHttpReques, Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma, FooBar, X-Custom-Header, X-XSRF-TOKEN, XSRF-TOKEN, Access-Control-Allow-Origin, strict-origin-when-cross-origin") // 允许跨域设置，可以返回其他子段
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                                  // 缓存请求信息，单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                         // 跨域请求是否需要带cookie信息，默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                                     // 设置返回格式是json
		}
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method
// 		// origin := c.Request.Header.Get("Origin")
// 		// var headerKeys []string
// 		// for k, _ := range c.Request.Header {
// 		// 	headerKeys = append(headerKeys, k)
// 		// }
// 		// headerStr := strings.Join(headerKeys, ", ")
// 		// if headerStr != "" {
// 		// 	headerStr = fmt.Sprintf("Access-Control-Allow-Origin, Acess-Control-Allow-Headers, %s", headerStr)
// 		// } else {
// 		// 	headerStr = "access-control-allow-origin, access-control-allow-headers"
// 		// }
		
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 			// c.Header("Access-Control-Allow-Origin", "*")
// 			c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
// 			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
// 			c.Writer.Header().Set("Access-Control-Max-Age", "172800")
// 			c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
// 			c.Set("content-type", "*")
		
// 		if method == "OPTIONS" {
// 			c.JSON(200, "Options Request!")
// 		}
// 		c.Next()
// 	}
// }
