package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	LOG   *zap.Logger
	CONF  *Config
	DB    *gorm.DB
	REDIS *redis.Client
)
