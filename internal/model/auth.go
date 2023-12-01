package model

import (
	"errors"
	"gorm.io/gorm"
)

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (auth Auth) TableName() string {
	return "blog_auth"
}

func (auth Auth) Get(db *gorm.DB) (Auth, error) {
	var a Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", auth.AppKey, auth.AppSecret, 0)
	err := db.First(&a).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return a, err
	}

	return a, nil
}
