package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var HPLogConfig = &LogConfig{
	PrintLog: true,
	Loglevel: "info", //默认级别
	// Filename:   filepath.Join(os.TempDir(), "HP-Log.log"),
	Filename:   filepath.Join("logs", "HP-Log.log"),
	MaxSize:    1024,  //日志文件大小
	MaxAge:     7,     //文件保存天数
	MaxBackups: 5,     //保存旧日志文件数量
	LocalTime:  true,  //分割的日志文件是否使用本地时间以及日志文件是否使用本地时间
	Compress:   false, // 是否压缩 disabled by default
	ModLevel: map[string]string{ //每隔模块的级别
		"main":   "info",
		"model1": "debug",
		"model2": "error",
	},
}

// LogConfig is the Configuration of log
type LogConfig struct {
	PrintLog   bool              `json:"printLog"`
	Loglevel   string            `json:"logLevel"`
	Filename   string            `json:"fileName"`
	MaxSize    int               `json:"maxSize"`
	MaxAge     int               `json:"maxAge"`
	MaxBackups int               `json:"maxBackups"`
	LocalTime  bool              `json:"localTime"`
	Compress   bool              `json:"compress"`
	ModLevel   map[string]string `json:"modLevel"`
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func LoadHPLogConfig(dir string, logConfig *LogConfig) (*LogConfig, error) {
	path := filepath.Join(dir, "./configs/log.toml")
	filePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	config := new(LogConfig)
	if CheckFileIsExist(filePath) { //文件存在
		if _, err := toml.DecodeFile(filePath, config); err != nil {
			return nil, err
		} else {
			HPLogConfig = config
		}
	} else {
		configBuf := new(bytes.Buffer)
		if logConfig != nil {
			if err := toml.NewEncoder(configBuf).Encode(logConfig); err != nil {
				return nil, err
			}
		} else {
			if err := toml.NewEncoder(configBuf).Encode(HPLogConfig); err != nil {
				return nil, err
			}
		}
		err := ioutil.WriteFile(filePath, configBuf.Bytes(), 0666)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
