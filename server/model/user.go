package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid"      gorm:"<-:create;not null;unique;index"` // could not modify
	Username string    `json:"user_name" gorm:"not null;index"`
	Password []byte    `json:"-"         gorm:"not null"`
	NickName string    `json:"nick_name" gorm:"not null"`
	Role     Role      `json:"role"      gorm:"default:0"`
}
