package middleware

import (
	"fmt" 
	"jwt-go-system/initializers"
	"jwt-go-system/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// 获取token
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
fmt.Println(tokenStr)
	// 验证token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if cliams, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 过期
		if float64(time.Now().Unix()) > cliams["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// 根据jwt携带消息(userId)查询数据库
		var user models.User
		initializers.DB.First(&user, cliams["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
