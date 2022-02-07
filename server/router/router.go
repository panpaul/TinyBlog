package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"server/global"
	_ "server/router/docs"
	"server/utils"
	"time"
)

func InitRouters() *gin.Engine {
	gin.SetMode(utils.If(global.CONF.Development, gin.DebugMode, gin.ReleaseMode).(string))

	r := gin.New()
	if global.CONF.Development {
		// Gin's default logger is pretty
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(ginzap.Ginzap(global.LOG, time.RFC3339, true))
		r.Use(ginzap.RecoveryWithZap(global.LOG, true))
	}
	r.MaxMultipartMemory = 8 << 20
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Cookie"},
		AllowCredentials: true,
	}))

	if global.CONF.Development {
		// add prof
		pprof.Register(r)
		// add swagger document
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// api v1
	apiV1 := r.Group("/api/v1")
	InitV1(apiV1)

	// static files
	r.Static("/static", "./resource/frontend/static")
	r.StaticFile("/", "./resource/frontend/index.html")
	r.StaticFile("/robots.txt", "./resource/frontend/robots.txt")
	r.NoRoute(func(c *gin.Context) {
		c.File("./resource/frontend/index.html")
	})
	return r
}
