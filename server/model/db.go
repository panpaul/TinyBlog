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
	"os"
	"server/global"
	"time"
)

func SetupDatabase() {
	global.LOG.Info("Connecting to database...")

	logger := zapgorm2.New(global.LOG)
	logger.IgnoreRecordNotFoundError = true

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
			Logger:      logger,
		})

	if err != nil {
		global.LOG.Error("Could not connect to database", zap.Error(err))
		return
	}

	checkAlive(3)
	migrateDatabase()

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

func migrateDatabase() {
	global.LOG.Info("Auto Migrating database...")

	_ = global.DB.AutoMigrate(&User{})
	_ = global.DB.AutoMigrate(&Article{})
}

func checkAlive(retry int) {
	if retry <= 0 {
		global.LOG.Fatal("All retries are used up. Failed to connect to database")
		os.Exit(-1)
	}

	db, err := global.DB.DB()
	if err != nil {
		global.LOG.Warn("failed to get sql.DB instance", zap.Error(err))
		time.Sleep(5 * time.Second)
		checkAlive(retry - 1)
		return
	}

	err = db.Ping()
	if err != nil {
		global.LOG.Warn("failed to ping database", zap.Error(err))
		time.Sleep(5 * time.Second)
		checkAlive(retry - 1)
		return
	}

	global.LOG.Info("database connect established")
}
