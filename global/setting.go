package global

import (
	"github.com/wow-unbelievable/blog/pkg/logger"
	"github.com/wow-unbelievable/blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettings
	EmailSetting    *setting.EmailSettingS
	TracerSetting   *setting.TraceSettingS
)
