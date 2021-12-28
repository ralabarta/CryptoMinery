package main

import (
	"github.com/IEatLemons/BaseGin/command"
	"github.com/IEatLemons/BaseGin/config"
	"github.com/IEatLemons/BaseGin/logger"
	"github.com/IEatLemons/GoHelper/helper"
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

	OVER := make(chan bool)
	_, err = command.ScheduledCommands(vHandle.Config.Scheduler)
	if err != nil {
		logger.Error("ScheduledCommands failed", zap.String("error", err.Error()))
		panic("init ScheduledCommands failed:" + err.Error())
	}
	<-OVER
}
