package middleware

import (
	"github.com/gin-gonic/gin"
	"server/e"
	"server/global"
	"server/service"
)

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		claim, err := service.JwtApp.ParseToken(token)
		if err != e.Success {
			global.Pong(err, nil, c)
			c.Abort()
			return
		}

		if !service.JwtApp.VersionValid(claim) {
			global.Pong(e.TokenRevoked, nil, c)
			c.Abort()
		}

		c.Set("claim", claim)
		c.Next()
	}
}
