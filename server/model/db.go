package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"server/global"
)

func SetupDatabase() {
	global.LOG.Info("Connecting to database...")

	var err error
	global.DB, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
				global.CONF.Database.User,
				global.CONF.Database.Password,
				global.CONF.Database.Database,
				global.CONF.Database.Address,
				global.CONF.Database.Port)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: global.CONF.Database.Prefix,
			},
			PrepareStmt: true,
			Logger:      zapgorm2.New(global.LOG),
		})
	if err != nil {
		global.LOG.Error("Could not connect to database", zap.Error(err))
		return
	}

	global.LOG.Info("Connecting to redis...")

	global.REDIS = redis.NewClient(&redis.Options{
		Addr:     global.CONF.Redis.Address,
		Password: global.CONF.Redis.Password,
		DB:       global.CONF.Redis.Db,
	})
	pong, err := global.REDIS.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Panic("redis connect error", zap.Error(err))
		return
	}
	global.LOG.Info("redis connect success", zap.String("pong", pong))
}

func MigrateDatabase() {
	global.LOG.Info("Migrating database...")

	_ = global.DB.AutoMigrate(User{})
}
