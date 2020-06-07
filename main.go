package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"switch-onchain/internal/public"

	srvhttp "switch-onchain/internal/server/http"

	"switch-onchain/internal/log"
	"switch-onchain/internal/server/profiling"
	"switch-onchain/internal/service"
	_ "switch-onchain/swagger"

	"github.com/devfabric/HP-Log/logger"
)

const ModuleName string = "main"

var (
	d string
	p string
	e bool
	v bool
)

func init() {
	flag.StringVar(&d, "d", "./", "set config dir path")
	flag.StringVar(&p, "p", "", "set port enabling profiling")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Usage = usage
}

func usage() {
	flag.PrintDefaults()
}

func main() {
	var (
		runDir  string
		err     error
		svc     *service.Service
		httpSrv *http.Server
	)

	flag.Parse()
	if v {
		fmt.Println(public.GetAppInfo())
		return
	}

	if d != "./" {
		runDir = d
	} else {
		runDir, err = public.GetCurrentDirectory()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	logConfig, err := log.LoadLogConfig(runDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	hpLog := logger.InitLog()
	logger := log.GetLogger(log.ModMain, hpLog)
	logger.Debug(*logConfig)

	if p != "" {
		err = profiling.CheckAndEnableProfiling(p)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if !public.CheckFileIsExist(filepath.Join(runDir, "configs")) {
		logger.Debug("configs dir not exist")
		return
	}

	for {
		svc, err = service.New(runDir, hpLog)
		if err != nil {
			logger.Errorf("new service err, %s", err.Error())
			goto end
		} else {
			logger.Debug("new service success")
		}

		httpSrv, err = srvhttp.New(runDir, svc)
		if err != nil {
			logger.Errorf("new http server err, %s", err.Error())
			goto end
		} else {
			logger.Debug("new http server success")
		}
		break

	end:
		if httpSrv != nil {
			httpSrv.Close()
		}
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logger.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				logger.Errorf("httpSrv.Shutdown error(%v)", err)
			}
			logger.Info("switch-onchain close...")
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
