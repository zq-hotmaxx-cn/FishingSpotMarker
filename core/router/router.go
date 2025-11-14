package router

import (
	"FishingSpotMarker/core/conf"
	"FishingSpotMarker/core/log"
	"FishingSpotMarker/pkg/config"
	"FishingSpotMarker/pkg/middleware"
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	fishing_spot_marker_router *gin.Engine
)

func Init() error {
	// 设置运行模式
	switch conf.GetConf().Server.Mode {
	case config.SERVER_MODE_DEV:
		gin.SetMode(gin.DebugMode)
	case config.SERVER_MODE_PROD:
		gin.SetMode(gin.ReleaseMode)
	default:
		return errors.New("Invalid server mode: " + conf.GetConf().Server.Mode)
	}

	// 创建路由
	fishing_spot_marker_router = gin.New()

	// 加载中间件
	loadMiddleware()

	// 加载路由
	loadRouter()

	return nil
}

func loadMiddleware() {
	fishing_spot_marker_router.
		Use(middleware.ZapLogger(log.GetZapLogger())).
		Use(middleware.StdoutLogger(log.GetStdInfoLogger())).
		Use(middleware.StderrLogger(log.GetStdErrorLogger())).
		Use(gin.Recovery()).
		Use(middleware.Cors())
}

func loadRouter() {
	loadApiRouter()
	loadAdminRouter()
}

func GetRouter() *gin.Engine {
	return fishing_spot_marker_router
}
