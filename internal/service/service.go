package service

import (
	"os"
	log "switch-onchain/internal/log"

	hplog "github.com/devfabric/HP-Log/logger"
	"github.com/devfabric/fabric-client/config"
	"github.com/devfabric/fabric-client/fabsdk"
	token "github.com/devfabric/trust-token/token"
)

var (
	logger *log.Logger
)

type Service struct {
	FabClient *fabsdk.FabricClient
	// DB        *dao.SqlxDB
	HPLog  *hplog.HPLog
	TKSrv  *token.TokenSrv
	IsExit chan bool
}

func New(dir string, hpLog *hplog.HPLog) (*Service, error) {
	s := &Service{
		HPLog: hpLog,
	}

	logger := log.GetLogger(log.ModService, s.HPLog)
	logger.Debug("switch-onchain start......")

	//init db
	//init fabric
	workDirForFabSDK := os.Getenv("WORKDIR")
	if workDirForFabSDK == "" {
		os.Setenv("WORKDIR", dir)
	}

	fabConfig, err := config.LoadFabircConfig(dir)
	if err != nil {
		logger.Debug("load fabirc config")
		return nil, err
	}
	s.FabClient = fabsdk.NewFabricClient(fabConfig.ConfigFile, fabConfig.ChannelID, fabConfig.UserName, fabConfig.UserOrg)
	err = s.FabClient.Setup(dir)
	if err != nil {
		logger.Errorf("fabric client setup err %s", err.Error())
		return nil, err
	}

	//init token服务对象
	tkSrv, err := token.GetTokenSrv(dir, "configs/token.toml")
	if err != nil {
		return nil, err
	}
	s.TKSrv = tkSrv

	s.IsExit = make(chan bool)
	return s, nil
}

func (s *Service) Close() {
	if s != nil {
		close(s.IsExit)
		// close db
		if s.FabClient != nil {
			s.FabClient.Close()
			s.FabClient = nil
		}
	}
}
