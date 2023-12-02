package model

import (
	"BloginGin/global"
	"BloginGin/pkg/setting"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

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
	s := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		databaseSetting.Host,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.DBName,
	)

	// 根据运行模式设置日志级别
	var logLevel logger.LogLevel
	if global.ServerSetting.RunMode == "debug" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{
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
	//db.Callback().Create().Replace("gorm:update_time_stamp",  BeforeUpdate)
	// 返回创建的数据库引擎实例
	return db, nil
}

//func BeforeCreate(tx *gorm.DB) (err error) {
//
//	return
//
//}
//
//func BeforeUpdate(tx *gorm.DB) error {
//	if _, ok := tx.Statement.Context.Value("gorm:update_column").(string); !ok {
//		_ = uint32(time.Now().Unix())
//	}
//	return nil
//}
//
//func (t Tag) BeforeDelete(tx *gorm.DB) (err error) {
//	var extraOption string
//	if str, ok := tx.Statement.Context.Value("gorm:delete_option").(string); !ok {
//		extraOption = fmt.Sprint(str)
//	}
//	deleteOnField, hasDeletedOnField := tx.Get("DeletedOn")
//	isDelFiled, hasIsDelField := tx.Get("IsDel")
//	if!hasDeletedOnField &&!hasIsDelField {
//		now := time.Now().Unix()
//		tx.
//	}
//
//	return
//}
//
//// 在查询记录之前调用的回调
//func (t Tag) AfterFind(tx *gorm.DB) (err error) {
//	// 在这里可以执行在查询记录之后需要进行的操作
//	return
//}
