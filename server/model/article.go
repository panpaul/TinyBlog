package model

import (
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UUID     uuid.UUID      `json:"uuid"      gorm:"<-:create;not null;unique;primaryKey"`
	AuthorID uuid.UUID      `json:"author_id" gorm:"<-:create;not null;index"`
	Author   User           `json:"author"    gorm:"foreignkey:AuthorID;references:UUID"`
	Title    string         `json:"title"     gorm:"not null"`
	Content  string         `json:"content"   gorm:"not null"`
	Level    Role           `json:"role"      gorm:"default:0"`
	Tags     pq.StringArray `json:"tags"      gorm:"type:text[]"`
}
