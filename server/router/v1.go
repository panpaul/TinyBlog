package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/global"
	"server/model"
	_ "server/router/docs"
	"server/service"
)

// InitV1 godoc
// @title                      TinyBlog API v1
// @version                    1.0.0
// @description                A Fresh New Blog System Written in Golang
// @contact.name               Paul
// @contact.url                https://blog.ofortune.xyz
// @contact.email              panyuxuan@hotmail.com
// @BasePath                   /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       token
func InitV1(v1g *gin.RouterGroup) {
	v1g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to TinyBlog API v1",
		})
	})

	for _, v := range EndpointsV1() {
		group := v1g.Group(v.Path)
		v.Register(group)
	}

}

func SetupRouterV1() {
	global.LOG.Info("Setting up router...")
	for _, v := range EndpointsV1() {
		v.Setup("/api/v1" + v.Path)
	}
	service.CasbinApp.Refresh()
}

func EndpointsV1() []model.EndpointInfo {
	return []model.EndpointInfo{
		{Path: "/test", Register: testApi, Setup: testSetup},
		{Path: "/user", Register: v1.UserApi, Setup: v1.UserSetup},
		{Path: "/article", Register: v1.ArticleApi, Setup: v1.ArticleSetup},
	}
}
