package controllers

import (
	resp "github.com/IEatLemons/GoHelper/helper/responses"
	"github.com/IEatLemons/GoHelper/middle"
	"github.com/gin-gonic/gin"
)

// BaseRoute /
func BaseRoute(g *gin.RouterGroup) {
	gin.SetMode(gin.ReleaseMode)
	M := middle.New()
	g.Use(M.Cors(), M.AuthLanguage(), M.AuthPlatform(""))
	resp.InitResp(M.Language)

	Base := g.Group("/")

	RegisteredRouter(Base)
}

// RegisteredRouter 注册相关模块的路由/
func RegisteredRouter(g *gin.RouterGroup) {

}
