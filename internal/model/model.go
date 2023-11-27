package model

import (
	"BloginGin/global"
	"BloginGin/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 共同具有的属性
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// NewDBEngine 根据给定的数据库设置创建一个新的 GORM 数据库引擎。
//
// 参数:
//   - databaseSetting: 包含数据库连接配置的结构体，包括用户名、密码、主机等信息。
//
// 返回值:
//   - *gorm.DB: GORM 数据库引擎实例。
//   - error: 如果创建过程中发生错误，则返回相应的错误信息。
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// 构建数据库连接字符串
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	// 根据运行模式设置日志级别
	var logLevel logger.LogLevel
	if global.ServerSetting.RunMode == "debug" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	// 获取 SQL.DB 对象，设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	// 返回创建的数据库引擎实例
	return db, nil
}
