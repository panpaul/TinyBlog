package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"server/e"
	"server/global"
	"server/middleware"
	"server/model"
	"server/service"
)

func testApi(c *gin.RouterGroup) {
	c.GET("/token", func(c *gin.Context) {
		claim := model.Claims{
			UUID:     uuid.UUID{},
			Username: "dev",
			NickName: "nick",
			Role:     model.RoleAdmin,
		}
		token, err := service.JwtApp.SignClaim(claim)
		global.Pong(err, token, c)
	})

	c.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	auth := c.Group("/auth")
	auth.Use(middleware.JwtHandler())
	auth.GET("/ping", func(c *gin.Context) {
		claim, exist := c.Get("claim")
		global.Pong(e.Success, gin.H{"claim": claim, "exist": exist}, c)
	})

	authorize := c.Group("/authorize")
	authorize.Use(middleware.JwtHandler())
	authorize.Use(middleware.CasbinHandler())

	authorize.GET("/ping", func(c *gin.Context) {
		claim, exist := c.Get("claim")
		global.Pong(e.Success, gin.H{"claim": claim, "exist": exist}, c)
	})
}

func testSetup(base string) {
	service.CasbinApp.AddRange([]model.CasbinRule{
		{
			Role:   model.RoleAdmin,
			Path:   base + "/authorize/*",
			Method: "(GET)|(POST)",
		},
	})
}
