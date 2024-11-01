package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 템플릿 파일 위치 설정
	r.LoadHTMLGlob("templates/*")

	// 기본 라우터 설정
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run()
}
