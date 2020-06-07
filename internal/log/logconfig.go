package log

import (
	"path/filepath"

	"github.com/devfabric/HP-Log/config"
)

const (
	ModMain    = "main"
	ModService = "service"
	ModHttp    = "http"
)

var HPLogConfig = &config.LogConfig{
	PrintLog:   true,
	Loglevel:   "info", //默认级别
	Filename:   filepath.Join("logs", "HP-Log.log"),
	MaxSize:    1024,  //日志文件大小
	MaxAge:     7,     //文件保存天数
	MaxBackups: 5,     //保存旧日志文件数量
	LocalTime:  true,  //分割的日志文件是否使用本地时间以及日志文件是否使用本地时间
	Compress:   false, // 是否压缩 disabled by default
	ModLevel: map[string]string{ //每隔模块的级别
		ModMain:    "debug",
		ModService: "debug",
	},
}

func LoadLogConfig(dir string) (*config.LogConfig, error) {
	logConfig, err := config.LoadHPLogConfig(dir, HPLogConfig)
	if err != nil {
		return nil, err
	}
	return logConfig, nil
}
