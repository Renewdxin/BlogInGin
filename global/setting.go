package global

import (
	"github.com/Renewdxin/BlogInGin/pkg/logger"
	"github.com/Renewdxin/BlogInGin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	Logger          *logger.Logger
)
