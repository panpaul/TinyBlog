package v1

import (
	"github.com/gin-gonic/gin"
	"server/e"
	"server/global"
	"server/middleware"
	"server/model"
	"server/service"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var l LoginForm
	_ = c.ShouldBindJSON(&l)
	u := &model.User{Username: l.Username, Password: []byte(l.Password)}
	if user, err := service.UserApp.Login(u); err != e.Success {
		global.Pong(err, nil, c)
	} else {
		tokenNext(c, user)
	}
}

func tokenNext(c *gin.Context, user *model.User) {
	claim := model.Claims{
		UUID:     user.UUID,
		Username: user.Username,
		NickName: user.NickName,
		Role:     user.Role,
	}
	token, err := service.JwtApp.SignClaim(claim)
	global.Pong(err, token, c)
}

func UserApi(c *gin.RouterGroup) {
	c.POST("/login", Login)
	c.GET("/user", middleware.JwtHandler(), func(ctx *gin.Context) {
		global.Pong(e.Success, "hello", ctx)
	})
}

func UserSetup(base string) {

}
