package model

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UUID     uuid.UUID `json:"uuid"`
	UserName string    `json:"user_name"`
	NickName string    `json:"nick_name"`
	Role     Role      `json:"role"`
	Version  int64     `json:"version"`
	jwt.RegisteredClaims
}
