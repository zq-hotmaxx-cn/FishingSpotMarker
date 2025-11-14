package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	std_error_prefix_text = "ERROR"
	std_info_prefix_text  = "INFO"
	std_warn_prefix_text  = "WARN"
	std_debug_prefix_text = "DEBUG"
)

func NewErrorLogger() *log.Logger {
	return log.New(os.Stderr, fmt.Sprintf("[%s] ", StdRedString(std_error_prefix_text)), log.LstdFlags)
}

func NewInfoLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] ", StdGreenString(std_info_prefix_text)), log.LstdFlags)
}

func NewWarnLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] ", StdYellowString(std_warn_prefix_text)), log.LstdFlags)
}

func NewDebugLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%s] ", StdBlueString(std_debug_prefix_text)), log.LstdFlags)
}
