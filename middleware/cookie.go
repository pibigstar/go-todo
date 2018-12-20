package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const TokenMaxAge = 86400

// RefreshTokenCookie 刷新过期时间
func RefreshTokenCookie(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	fmt.Println(err)
	if tokenString != "" && err == nil {
		c.SetCookie("token", tokenString, TokenMaxAge, "/", "", true, true)
		// if user, err := getUser(c); err == nil {
		// 	model.UserToRedis(user)
		// }
	}
	c.Next()
}
