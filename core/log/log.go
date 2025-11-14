package log

import (
	"FishingSpotMarker/core/conf"
	"FishingSpotMarker/pkg/logger"
)

func Init() error {
	// 创建 std logger
	fishing_spot_marker_std_error_logger = logger.NewErrorLogger()
	fishing_spot_marker_std_info_logger = logger.NewInfoLogger()
	fishing_spot_marker_std_warn_logger = logger.NewWarnLogger()
	fishing_spot_marker_std_debug_logger = logger.NewDebugLogger()

	// 创建 zap logger
	zap_logger, err := logger.LoadZapLogger(conf.GetConf().Server.Log)
	if err != nil {
		return err
	}
	fishing_spot_marker_zap_logger = zap_logger

	return nil
}
