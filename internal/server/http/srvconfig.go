package http

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"switch-onchain/internal/public"

	"github.com/BurntSushi/toml"
)

var (
	cfg  *HttpServer
	once sync.Once
	lock sync.RWMutex
)

const (
	SRV_PORT          = 8816
	SRV_ADDRESS       = "0.0.0.0"
	SRV_MIDWARE       = 1
	SRV_READ_TIMEOUT  = 30
	SRV_WRITE_TIMEOUT = 30
	SRV_TLS_ENABLED   = false
	SRV_TLS_CERTFILE  = "app_client.pem"
	SRV_TLS_KEY       = "app_client.key"
	SRV_TLS_CACERT    = "app_ca.pem"
)

type ServerTLSConfig struct {
	Enabled  bool   `toml:"enabled"`
	CertFile string `toml:"certfile"`
	KeyFile  string `toml:"keyfile"`
	CACert   string `toml:"cacert"`
}

type HttpServer struct {
	Port         int             `toml:"port"`
	Address      string          `toml:"address"`
	MidWare      int             `toml:"midware"`
	ReadTimeout  int             `toml:"readtimeout"`
	WriteTimeout int             `toml:"writetimeout"`
	TLS          ServerTLSConfig `toml:"tls"`
}

func Config(dir string) *HttpServer {
	once.Do(func() {
		ReloadConfig(dir)
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT) //kill -SIGINT 6666
		go func() {
			for {
				s := <-c
				switch s {
				case syscall.SIGINT:
					ReloadConfig(dir)
				}
			}
		}()
	})

	lock.Lock()
	defer lock.Unlock()
	return cfg
}

func ReloadConfig(dir string) {
	path := filepath.Join(dir, "./configs/http.toml")
	filePath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	config := new(HttpServer)
	if public.CheckFileIsExist(filePath) { //文件存在
		if _, err := toml.DecodeFile(filePath, config); err != nil {
			panic(err)
		}

		logger.Debug("http server config ", config)
		lock.Lock()
		defer lock.Unlock()
		cfg = config
	} else {
		config.Port = SRV_PORT
		config.Address = SRV_ADDRESS
		config.MidWare = SRV_MIDWARE
		config.ReadTimeout = SRV_READ_TIMEOUT
		config.WriteTimeout = SRV_WRITE_TIMEOUT
		config.TLS.Enabled = SRV_TLS_ENABLED
		config.TLS.CertFile = SRV_TLS_CERTFILE
		config.TLS.KeyFile = SRV_TLS_KEY
		config.TLS.CACert = SRV_TLS_CACERT

		configBuf := new(bytes.Buffer)
		if err := toml.NewEncoder(configBuf).Encode(config); err != nil {
			logger.Errorf("NewEncoder Err %s ", err.Error())
			return
		}
		err := ioutil.WriteFile(filePath, configBuf.Bytes(), 0666)
		if err != nil {
			logger.Errorf("WriteFile Err %s ", err.Error())
			return
		}
		cfg = config
	}
}
