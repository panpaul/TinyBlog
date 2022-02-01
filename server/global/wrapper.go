package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/e"
	"server/utils"
)

func Wrap(err e.Err, body interface{}) interface{} {
	return Response{
		Code: int(err),
		Msg:  err.String(),
		Body: utils.If(err == e.Success, body, nil),
	}
}

func Pong(err e.Err, body interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Wrap(err, body))
}
