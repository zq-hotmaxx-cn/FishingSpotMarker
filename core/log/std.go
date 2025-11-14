package log

import (
	"FishingSpotMarker/core/conf"
	"FishingSpotMarker/pkg/config"
	"FishingSpotMarker/pkg/logger"
	"log"
)

var (
	fishing_spot_marker_std_error_logger = logger.NewErrorLogger()
	fishing_spot_marker_std_info_logger  = logger.NewInfoLogger()
	fishing_spot_marker_std_warn_logger  = logger.NewWarnLogger()
	fishing_spot_marker_std_debug_logger = logger.NewDebugLogger()
)

func GetStdErrorLogger() *log.Logger {
	return fishing_spot_marker_std_error_logger
}

func GetStdInfoLogger() *log.Logger {
	return fishing_spot_marker_std_info_logger
}

func GetStdWarnLogger() *log.Logger {
	return fishing_spot_marker_std_warn_logger
}

func GetStdDebugLogger() *log.Logger {
	return fishing_spot_marker_std_debug_logger
}

func StdErrorf(format string, v ...any) {
	fishing_spot_marker_std_error_logger.Printf(format, v...)
}

func StdInfof(format string, v ...any) {
	fishing_spot_marker_std_info_logger.Printf(format, v...)
}

func StdWarnf(format string, v ...any) {
	fishing_spot_marker_std_warn_logger.Printf(format, v...)
}

func StdDebugf(format string, v ...any) {
	if conf.GetConf().Server.Mode == config.SERVER_MODE_DEV {
		fishing_spot_marker_std_debug_logger.Printf(format, v...)
	}
}
