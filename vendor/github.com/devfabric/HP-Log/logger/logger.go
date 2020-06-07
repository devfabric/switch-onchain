package logger

import (
	"os"
	"sync"
	"time"

	"github.com/devfabric/HP-Log/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type HPLog struct {
	LogMap   map[string]*zap.Logger
	ZapWrite zapcore.WriteSyncer
	Encoder  zapcore.Encoder
	Mutex    sync.Mutex
}

var (
	LogLevelMap = map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"error": zap.ErrorLevel,
	}

	GetLogLevel = func(logLevel string) zapcore.Level {
		var (
			level zapcore.Level
			ok    bool
		)
		if level, ok = LogLevelMap[logLevel]; ok {
			return level
		}
		return zap.InfoLevel
	}
)

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	enc.AppendString(t.Format(time.RFC3339Nano))
}

func InitLog() *HPLog {
	hook := lumberjack.Logger{
		Filename:   config.HPLogConfig.Filename,   // 日志文件路径
		MaxSize:    config.HPLogConfig.MaxSize,    // megabytes
		MaxBackups: config.HPLogConfig.MaxBackups, // 最多保留3个备份
		MaxAge:     config.HPLogConfig.MaxAge,     //days
		LocalTime:  config.HPLogConfig.LocalTime,
		Compress:   config.HPLogConfig.Compress, // 是否压缩 disabled by default
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	if hook.LocalTime {
		encoderConfig.EncodeTime = TimeEncoder //本地定义化时间格式
	} else {
		encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder //ISO8601TimeEncoder
	}

	hpLog := &HPLog{
		LogMap: make(map[string]*zap.Logger),
	}

	hpLog.ZapWrite = zapcore.AddSync(&hook)
	hpLog.Encoder = zapcore.NewConsoleEncoder(encoderConfig)

	return hpLog
}

func InitLogger(module string) (*HPLog, *zap.SugaredLogger) {
	hook := lumberjack.Logger{
		Filename:   config.HPLogConfig.Filename,   // 日志文件路径
		MaxSize:    config.HPLogConfig.MaxSize,    // megabytes
		MaxBackups: config.HPLogConfig.MaxBackups, // 最多保留3个备份
		MaxAge:     config.HPLogConfig.MaxAge,     //days
		LocalTime:  config.HPLogConfig.LocalTime,
		Compress:   config.HPLogConfig.Compress, // 是否压缩 disabled by default
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	if hook.LocalTime {
		encoderConfig.EncodeTime = TimeEncoder //本地定义化时间格式
	} else {
		encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder //ISO8601TimeEncoder
	}

	var level zapcore.Level
	hpLog := &HPLog{
		LogMap: make(map[string]*zap.Logger),
	}

	hpLog.ZapWrite = zapcore.AddSync(&hook)
	hpLog.Encoder = zapcore.NewConsoleEncoder(encoderConfig)
	if logLevel, ok := config.HPLogConfig.ModLevel[module]; ok {
		level = GetLogLevel(logLevel)
	} else {
		switch config.HPLogConfig.Loglevel {
		case "debug":
			level = zap.DebugLevel
		case "info":
			level = zap.InfoLevel
		case "error":
			level = zap.ErrorLevel
		default:
			level = zap.InfoLevel
		}
	}

	var zapCore zapcore.Core
	if config.HPLogConfig.PrintLog {
		zapCore = zapcore.NewCore(
			hpLog.Encoder, // 编码器配置
			// hpLog.ZapWrite,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), hpLog.ZapWrite), // 打印到控制台和文件
			level,
		)
	} else {
		zapCore = zapcore.NewCore(
			hpLog.Encoder, // 编码器配置
			hpLog.ZapWrite,
			level,
		)
	}

	hpLog.LogMap[module] = zap.New(zapCore)
	hpLog.LogMap[module].Info("InitLogger Init Success")

	return hpLog, hpLog.LogMap[module].Sugar()
}

func (h *HPLog) GetLogger(module string) *zap.SugaredLogger {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	hpLog, ok := h.LogMap[module]
	if ok {
		return hpLog.Sugar()
	}

	var level zapcore.Level
	if logLevel, ok := config.HPLogConfig.ModLevel[module]; ok {
		level = GetLogLevel(logLevel)
	} else {
		switch config.HPLogConfig.Loglevel {
		case "debug":
			level = zap.DebugLevel
		case "info":
			level = zap.InfoLevel
		case "error":
			level = zap.ErrorLevel
		default:
			level = zap.InfoLevel
		}
	}

	var zapCore zapcore.Core
	if config.HPLogConfig.PrintLog {
		zapCore = zapcore.NewCore(
			h.Encoder, // 编码器配置
			// hpLog.ZapWrite,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), h.ZapWrite), // 打印到控制台和文件
			level,
		)
	} else {
		zapCore = zapcore.NewCore(
			h.Encoder, // 编码器配置
			h.ZapWrite,
			level,
		)
	}

	// zapCore := zapcore.NewCore(
	// 	h.Encoder,
	// 	h.ZapWrite,
	// 	level,
	// )

	h.LogMap[module] = zap.New(zapCore)
	return h.LogMap[module].Sugar()
}
