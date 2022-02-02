package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type Claims struct {
	UUID     uuid.UUID `json:"uuid"`
	UserName string    `json:"user_name"`
	NickName string    `json:"nick_name"`
	Role     Role      `json:"role"`
	Version  int64     `json:"version"`
	jwt.StandardClaims
}
