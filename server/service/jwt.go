package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"server/e"
	"server/global"
	"server/model"
	"time"
)

type JwtService struct {
	SigningKey []byte
}

var JwtApp = new(JwtService)

func (j *JwtService) Setup() {
	j.SigningKey = []byte(global.CONF.Jwt.JwtSecret)
}

func (j *JwtService) ParseToken(tokenText string) (*model.Claims, e.Err) {
	if tokenText == "" {
		return nil, e.TokenEmpty
	}

	token, err := jwt.ParseWithClaims(
		tokenText,
		&model.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil,
					fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return j.SigningKey, nil
		})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, e.TokenMalformed
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, e.TokenTimeError
		} else {
			global.LOG.Info("JWT Token Validation Error", zap.Error(err))
			return nil, e.InternalError
		}
	} else if err != nil {
		global.LOG.Warn("JWT Token Parse Error", zap.Error(err))
		return nil, e.InternalError
	}

	if token.Valid {
		c := token.Claims.(*model.Claims)
		return c, e.Success
	}

	return nil, e.InvalidParameter
}

func (j *JwtService) SignClaim(claim model.Claims) (string, e.Err) {
	now := time.Now()
	claim.IssuedAt = jwt.NewNumericDate(now)
	claim.NotBefore = jwt.NewNumericDate(now)
	claim.ExpiresAt = jwt.NewNumericDate(now.Add(time.Duration(global.CONF.Jwt.JwtExpireHour) * time.Hour))
	claim.Version = UserApp.NextVersion(claim.UUID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	ss, err := token.SignedString(j.SigningKey)
	if err != nil {
		global.LOG.Warn("jwt.SignedString error", zap.Error(err))
		return "", e.TokenSignError
	}
	return ss, e.Success
}

func (j *JwtService) VersionValid(claim *model.Claims) bool {
	return claim.Version == UserApp.GetVersion(claim.UUID)
}
