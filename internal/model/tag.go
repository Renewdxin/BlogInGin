package model

import (
	"BloginGin/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (tag Tag) TableName() string {
	return "blog_tag"
}

func (tag Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if tag.Name != "" {
		db = db.Where("name = ?", tag.Name)
	}
	db = db.Where("state = ?", tag.State)
	err := db.Model(&tag).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (tag Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if tag.Name != "" {
		db = db.Where("name = ?", tag.Name)
	}
	db = db.Where("state = ?", tag.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (tag Tag) Create(db *gorm.DB) error {
	return db.Create(&tag).Error
}

func (tag Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(tag).Where("id = ? AND is_del = ?", tag.ID).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (tag Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND id_del = ?", tag.ID, 0).Delete(&tag).Error
}
