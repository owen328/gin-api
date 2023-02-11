package middleware

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/model"
	"go_learn/common"
	"go_learn/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenParts := strings.SplitN(tokenString, " ", 2)
		if !(tokenParts[0] == "Bearer" && len(tokenParts) == 2) {
			utils.UnAuthorization(c)
			c.Abort()
			return
		}

		token, claims, err := utils.ParseToken(tokenParts[1])
		if err != nil || !token.Valid {
			utils.UnAuthorization(c)
			c.Abort()
			return
		}
		var user model.Admin
		if first := common.DB.Where("id=?", claims.UserId).First(&user); first.Error != nil {
			utils.UnAuthorization(c)
			c.Abort()
			return
		}
		c.Set("user_id", user.Id)
		c.Set("username", user.Username)
		c.Next()
	}
}