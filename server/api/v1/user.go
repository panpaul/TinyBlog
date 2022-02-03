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
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

// Login godoc
// @Summary     user login
// @Description accept username and password(plaintext) and return token
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       login body v1.UserForm true "username and password"
// @Router      /user/login [post]
func Login(c *gin.Context) {
	var l UserForm
	_ = c.ShouldBindJSON(&l)
	u := &model.User{UserName: l.UserName, Password: []byte(l.Password)}
	if user, err := service.UserApp.Login(u); err != e.Success {
		global.Pong(err, nil, c)
	} else {
		tokenNext(c, user)
	}
}

// Register godoc
// @Summary     user register
// @Description accept username and password(plaintext) and nickname(optionally) and return token
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       register body v1.UserForm true "username, nickname and password"
// @Router      /user/register [post]
func Register(c *gin.Context) {
	var r UserForm
	_ = c.ShouldBindJSON(&r)
	u := &model.User{
		UserName: r.UserName,
		Password: []byte(r.Password),
		NickName: utils.If(r.NickName == "", r.UserName, r.NickName).(string),
		Role:     model.RoleUser,
	}
	if user, err := service.UserApp.Register(u); err != e.Success {
		global.Pong(err, nil, c)
	} else {
		tokenNext(c, user)
	}
}

// Logout godoc
// @Summary     user logout
// @Description logout the user which described by the token
// @Tags        User
// @Produce     json
// @Router      /user/logout [post]
// @Security    ApiKeyAuth
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

// Profile godoc
// @Summary     user profile update
// @Description update nickname and(or) password
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       profile body v1.UserForm true "nickname and password"
// @Router      /user/profile [post]
// @Security    ApiKeyAuth
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
		UserName: user.UserName,
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
	global.LOG.Debug("api v1 /user setup")
}
