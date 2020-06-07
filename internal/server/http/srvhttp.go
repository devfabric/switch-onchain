package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	log "switch-onchain/internal/log"

	"strconv"
	"time"

	"switch-onchain/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	svc    *service.Service
	logger *log.Logger
)

type HttpSrv struct {
	Router    *gin.Engine
	Server    *http.Server
	SrvConfig *HttpServer
}

func NewHttpSrv(cfg *HttpServer) *HttpSrv {
	return &HttpSrv{
		Router:    nil,
		Server:    nil,
		SrvConfig: cfg,
	}
}

func New(dir string, s *service.Service) (*http.Server, error) {
	{
		logger = log.GetLogger(log.ModHttp, s.HPLog)
		logger.Debugf("%s", "Initializing http server")
	}

	//read config file
	config := Config(dir)

	svc = s
	httpSrv := NewHttpSrv(config)
	initRouter(svc, httpSrv)
	if err := httpSrv.Start(); err != nil {
		return nil, err
	}
	return httpSrv.Server, nil
}

func (hs *HttpSrv) Start() error {
	var (
		listener net.Listener
		exitErr  error
		addrStr  string
		addr     string
	)

	addr = net.JoinHostPort(hs.SrvConfig.Address, strconv.Itoa(hs.SrvConfig.Port))
	if hs.SrvConfig.TLS.Enabled {
		addrStr = fmt.Sprintf("https://%s", addr)
		pool := x509.NewCertPool()
		caCert, err := ioutil.ReadFile(hs.SrvConfig.TLS.CACert)
		if err != nil {
			logger.Errorf("read file %s, err %s\n", hs.SrvConfig.TLS.CACert, err.Error())
			return err
		}

		cert, err := tls.LoadX509KeyPair(hs.SrvConfig.TLS.CertFile, hs.SrvConfig.TLS.KeyFile)
		if err != nil {
			logger.Errorf(err.Error())
			return err
		}

		var config *tls.Config
		if caCert != nil {
			pool.AppendCertsFromPEM(caCert)
			config = &tls.Config{
				ClientCAs:  pool,
				ClientAuth: tls.RequireAndVerifyClientCert,

				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS10,
				MaxVersion:   tls.VersionTLS12,
			}
		} else {
			config = &tls.Config{
				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS10,
				MaxVersion:   tls.VersionTLS12,
			}
		}

		listener, exitErr = tls.Listen("tcp", addr, config)
		if err != nil {
			return fmt.Errorf("TLS listen failed for %s: %s", addrStr, exitErr)
		}
	} else {
		addrStr = fmt.Sprintf("http://%s", addr)
		listener, exitErr = net.Listen("tcp", addr)
		if exitErr != nil {
			return fmt.Errorf("TCP listen failed for %s: %s", addrStr, exitErr)
		}
	}

	hs.Server = &http.Server{
		Handler:        hs.Router,
		ReadTimeout:    time.Duration(hs.SrvConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(hs.SrvConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 10,
	}

	go func() {
		exitErr = hs.Server.Serve(listener)
		if errors.Cause(exitErr) == http.ErrServerClosed {
			logger.Info("gin httpserver closed")
			return
		}
		panic(errors.Wrapf(exitErr, "gin: engine.ListenServer(%+v, %+v)", hs.Server, listener))
	}()

	return nil
}

func initRouter(svc *service.Service, srv *HttpSrv) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	//swagger文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//组1: echo路由检测组与NoAuthHandle中间件
	// apiRoot := router.Group("/")
	// apiRoot.Use(srv.MW.NoAuthHandle)
	// {
	// 	apiRoot.GET("/", GET_Echo)
	// }

	//组2: 无路由限制与UnLimitHandle中间件
	// apiTokenG1 := router.Group("/token/v1")
	// apiTokenG1.Use(srv.MW.UnLimitHandle)
	// {
	// 	t := &Token{
	// 		Aes: &AESCache{
	// 			// Key: []byte(svc.UCenter.Config.GetAESKey()),
	// 			// IV:  []byte(svc.UCenter.Config.GetAESIV()),
	// 		},
	// 		SVC: svc,
	// 	}
	// 	apiTokenG1.POST("/generate", t.Generate)
	// 	apiTokenG1.POST("/verify", t.Verify)
	// }

	//组3: api路由与MWHandle中间件
	// apiGroupV1 := router.Group("/api/v1")
	// apiGroupV1.Use(srv.MW.MWHandle)
	// {
	// 	apiV1 := &v1.APIV1{
	// 		// CacheType: common.CacheTypeMap[srv.SrvConfig.Cache.CacheType],
	// 		SVC: svc,
	// 	}
	// 	_ = apiV1
	// 	//resauth module
	// 	apiGroupV1.POST("/app/api", apiV1.API)
	// 	// apiGroupV1.POST("/res/EnrollResAuthTest", apiV1.EnrollResAuthTest)
	// 	// apiGroupV1.POST("/res/grantResAuth", apiV1.GrantResAuth)
	// 	// apiGroupV1.POST("/res/updateResAuth", apiV1.UpdateResAuth)
	// 	// apiGroupV1.POST("/res/revokeResAuth", apiV1.RevokeResAuth)
	// 	// apiGroupV1.POST("/res/userResAuthList", apiV1.GetUserResAuthList)
	// 	// apiGroupV1.POST("/res/grantResAuthList", apiV1.GetGrantResAuthList)
	// 	// apiGroupV1.POST("/res/verify", apiV1.VerifyResAuth)
	// 	// apiGroupV1.POST("/res/deleteResAuth", apiV1.DeleteResAuth)
	// }

	srv.Router = router
}
