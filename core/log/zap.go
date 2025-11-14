package log

import (
	"FishingSpotMarker/pkg/logger"

	"go.uber.org/zap"
)

var (
	fishing_spot_marker_zap_logger *zap.Logger
)

func GetZapLogger() *zap.Logger {
	return fishing_spot_marker_zap_logger
}

func GetZapGormLogger() *logger.ZapGormLogger {
	return logger.NewZapGormLogger(fishing_spot_marker_zap_logger)
}
