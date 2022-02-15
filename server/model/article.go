package model

import (
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	UUID        uuid.UUID      `json:"uuid"        gorm:"<-:create;not null;unique;primaryKey"`
	AuthorID    uuid.UUID      `json:"-"           gorm:"<-:create;not null;index"`
	Author      User           `json:"author"      gorm:"foreignkey:AuthorID;references:UUID"`
	Title       string         `json:"title"       gorm:"not null"`
	Description string         `json:"description" gorm:""`
	Content     string         `json:"content"     gorm:"not null"`
	EnableMath  bool           `json:"enable_math" gorm:"not null;default:false"`
	Level       Role           `json:"-"           gorm:"default:0"`
	Tags        pq.StringArray `json:"tags"        gorm:"type:text[]"`
}

type ArticleDescription struct {
	UUID        uuid.UUID `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
