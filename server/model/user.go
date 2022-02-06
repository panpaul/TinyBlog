package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"-"         gorm:"<-:create;not null;unique;primaryKey"` // could not modify
	UserName   string    `json:"user_name" gorm:"not null;index"`
	Password   []byte    `json:"-"         gorm:"not null"`
	NickName   string    `json:"nick_name" gorm:"not null"`
	Role       Role      `json:"-"         gorm:"default:0"`
}
