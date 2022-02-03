package model

import (
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID      `json:"uuid"      gorm:"<-:create;not null;unique;primaryKey"`
	AuthorID   uuid.UUID      `json:"-"         gorm:"<-:create;not null;index"`
	Author     User           `json:"author"    gorm:"foreignkey:AuthorID;references:UUID"`
	Title      string         `json:"title"     gorm:"not null"`
	Content    string         `json:"content"   gorm:"not null"`
	Level      Role           `json:"-"         gorm:"default:0"`
	Tags       pq.StringArray `json:"tags"      gorm:"type:text[]"`
}
