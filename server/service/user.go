package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"server/e"
	"server/global"
	"server/model"
	"server/utils"
)

type UserService struct{}

var UserApp = new(UserService)

func (u *UserService) Register(user model.User) (model.User, e.Err) {
	var buf model.User
	if !errors.Is(
		global.DB.Where("username = ?", user.Username).First(&buf).Error,
		gorm.ErrRecordNotFound,
	) {
		return model.User{}, e.UserDuplicated
	}

	hashed, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		global.LOG.Warn("bcrypt.GenerateFromPassword error", zap.Error(err))
		return model.User{}, e.InternalError
	}

	user.UUID, _ = uuid.NewV4()
	user.Password = hashed

	err = global.DB.Create(&user).Error
	return user, utils.If(err == nil, e.Success, e.DBCreateError).(e.Err)
}

func (u *UserService) Login(user *model.User) (*model.User, e.Err) {
	plainPassword := user.Password

	if errors.Is(
		global.DB.Where("username = ?", user.Username).First(&user).Error,
		gorm.ErrRecordNotFound,
	) {
		return &model.User{}, e.UserNotFound
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, plainPassword); err != nil {
		return user, e.UserPasswordError
	}

	return user, e.Success
}

func (u *UserService) NextVersion(uuid uuid.UUID) int64 {
	val, err := global.REDIS.Incr(context.Background(), fmt.Sprintf("VER:%s", uuid.String())).Result()
	if err != nil {
		global.LOG.Warn("redis.Incr error", zap.Error(err))
		return -1
	}
	return val
}

func (u *UserService) GetVersion(uuid uuid.UUID) int64 {
	val, err := global.REDIS.Get(context.Background(), fmt.Sprintf("VER:%s", uuid.String())).Int64()
	if err != nil {
		global.LOG.Warn("redis.Get error", zap.Error(err))
		return -1
	}
	return val
}
