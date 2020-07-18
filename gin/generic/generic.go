package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) { //路由请求的泛绑定，所有/user/前缀的请求都会达到一个回调函数里
		c.String(200, "hello word") //返回的内容
	})
	r.Run(":80") //运行80口
}
