package model

import "github.com/gin-gonic/gin"

type EndpointInfo struct {
	Path     string
	Register func(group *gin.RouterGroup)
	Setup    func(base string)
}
