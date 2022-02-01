package v1

import (
	"github.com/gin-gonic/gin"
	"server/e"
	"server/global"
	"server/middleware"
	"server/model"
	"server/service"
	"server/utils"
)

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

func Login(c *gin.Context) {
	var l UserForm
	_ = c.ShouldBindJSON(&l)
	u := &model.User{Username: l.Username, Password: []byte(l.Password)}
	if user, err := service.UserApp.Login(u); err != e.Success {
		global.Pong(err, nil, c)
	} else {
		tokenNext(c, user)
	}
}

func Register(c *gin.Context) {
	var r UserForm
	_ = c.ShouldBindJSON(&r)
	u := &model.User{
		Username: r.Username,
		Password: []byte(r.Password),
		NickName: utils.If(r.NickName == "", r.Username, r.NickName).(string),
		Role:     model.RoleUser,
	}
	if user, err := service.UserApp.Register(u); err != e.Success {
		global.Pong(err, nil, c)
	} else {
		tokenNext(c, user)
	}
}

func Logout(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		global.Pong(e.Unauthenticated, nil, c)
		return
	}

	if service.UserApp.NextVersion(claim.(*model.Claims).UUID) == -1 {
		global.Pong(e.RedisIncrError, nil, c)
		return
	}

	global.Pong(e.Success, nil, c)
}

func Profile(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		global.Pong(e.Unauthenticated, nil, c)
		return
	}

	var form UserForm
	_ = c.ShouldBindJSON(&form)

	u := &model.User{
		UUID:     claim.(*model.Claims).UUID,
		NickName: form.NickName,
		Password: utils.If(form.Password != "", []byte(form.Password), nil).([]byte),
	}

	if user, err := service.UserApp.Update(u); err != e.Success {
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
	c.POST("/register", Register)
	c.POST("/logout", middleware.JwtHandler(), Logout)
	c.POST("/profile", middleware.JwtHandler(), Profile)
}

func UserSetup(base string) {

}
