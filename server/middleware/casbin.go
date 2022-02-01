package middleware

import (
	"github.com/gin-gonic/gin"
	"server/e"
	"server/global"
	"server/model"
	"server/service"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, exist := c.Get("claim")
		if !exist {
			global.Pong(e.Unauthenticated, nil, c)
			c.Abort()
			return
		}
		obj := c.Request.URL.Path
		act := c.Request.Method
		sub := claim.(*model.Claims).Role.String()
		enf := service.CasbinApp.Casbin()
		success, _ := enf.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			global.Pong(e.InsufficientPermission, nil, c)
			c.Abort()
			return
		}
	}
}
