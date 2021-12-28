package main

import (
	"fmt"

	"github.com/IEatLemons/BaseGin/api"
	"github.com/IEatLemons/BaseGin/config"
	"github.com/IEatLemons/BaseGin/link/base"
	"github.com/IEatLemons/BaseGin/logger"
	"github.com/IEatLemons/GoHelper/helper"
	"github.com/IEatLemons/GoHelper/middle"
	"github.com/IEatLemons/GoHelper/signal"
	"go.uber.org/zap"
)

func main() {
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	err := vHandle.InitConfig(path)
	if err != nil {
		logger.Error("init config failed", zap.String("error", err.Error()))
		panic("init config failed:" + err.Error())
	}
	// init logger
	if err := logger.InitLogger(&vHandle.Config.Log); err != nil {
		logger.Error("init logger failed", zap.String("error", err.Error()))
		panic("init logger failed:" + err.Error())
	}

	// init MySQL
	if err := base.InitEngineGroup(&vHandle.Config.MysqlMaster, &vHandle.Config.MysqlSlave); err != nil {
		logger.Error("init mysql failed", zap.String("error", err.Error()))
		panic("init mysql failed:" + err.Error())
	}
	// init redis
	if err := base.InitRedis(&vHandle.Config.Redis); err != nil {
		logger.Error("init redis failed", zap.String("error", err.Error()))
		panic("init redis failed:" + err.Error())
	}

	// init Middle
	middle.Init(base.Redis, vHandle.Config.InternalHost, vHandle.Config.InternalIp)

	// init router
	r := api.SetupApi()
	// listen and signal
	if err := signal.Listen(fmt.Sprintf(":%d", vHandle.Config.Server.Port), r); err != nil {
		logger.Error("listen and signal failed", zap.String("error", err.Error()))
		panic("init redis failed:" + err.Error())
	}

}
