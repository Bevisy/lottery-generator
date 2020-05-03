package main

import (
	"github.com/gin-gonic/gin"
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

	r.Run("127.0.0.1:8080")
}
