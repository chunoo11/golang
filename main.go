package main

import (
	"my-gin-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 데이터베이스에 User 모델을 마이그레이션
	db.AutoMigrate(&models.User{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 기본 페이지 라우트 설정
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 회원가입 페이지 렌더링
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	// 회원가입 폼 데이터 처리
	r.POST("/signup", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 비밀번호 암호화는 추후 구현 가능 (예: bcrypt)
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	})

	// 서버 실행
	r.Run()
}
