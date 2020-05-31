package main

import (
	"github.com/gin-gonic/gin"
	"lottery/common"
	"lottery/generater"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 全局初始化随机数 seed

	r := gin.Default()
	r.GET("/lottery", func(c *gin.Context) {
		balls := generater.DoubleColorBall()

		c.String(http.StatusOK, strings.Join(balls, " "))
	})

	r.GET("/help", func(c *gin.Context) {
		c.String(http.StatusOK, common.Help())
	})

	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/help" //重定向
		r.HandleContext(c)
	})

	r.GET("/blog", func(c *gin.Context) {
		//重定向
		c.Redirect(http.StatusMovedPermanently, "https://bevisy.github.io/")
	})

	r.Run("0.0.0.0:80")
}
