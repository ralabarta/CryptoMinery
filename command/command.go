package command

import (
	"github.com/IEatLemons/BaseGin/config"
	"github.com/IEatLemons/BaseGin/logger"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"strings"
)

// ScheduledCommands 定时任务
func ScheduledCommands(s config.Scheduler) (c *cron.Cron, err error) {
	c = cron.New()
	for ClassMethod, timing := range s {
		ClassMethodSlice := strings.Split(ClassMethod, "/")
		if len(ClassMethodSlice) != 1 {
			continue
		}
		method := ClassMethodSlice[0]
		var function func()
		function = methodMap(method)
		if function != nil {
			if err := c.AddFunc(timing, function); err != nil {
				logger.Error("scheduled tasks failed", zap.String("error", err.Error()))
			}
		}
	}
	c.Start()
	return
}

// order 订单服务
func methodMap(method string) (function func()) {
	switch method {

	}
	return
}
