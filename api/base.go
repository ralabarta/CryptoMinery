package api

import (
	"github.com/IEatLemons/BaseGin/business/controllers"
	"github.com/IEatLemons/BaseGin/logger"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func SetupApi() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	r.SetFuncMap(template.FuncMap{
		"timeStr": func(timestamp int64) string {
			return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
		},
	})

	g := r.Group("/")

	controllers.BaseRoute(g)

	return r
}
