package helpers

import (
	"log"
	"os"
)

type level int8

const (
	PANIC_LEVEL level = iota
	ERROR_LEVEL
	WARN_LEVEL
	INFO_LEVEL
	DEBUG_LEVEL
	TRACE_LEVEL
)

func GetLevel(verbosity int) level {
	switch verbosity {
	case 0:
		return PANIC_LEVEL
	case 1:
		return ERROR_LEVEL
	case 2:
		return WARN_LEVEL
	case 3:
		return INFO_LEVEL
	case 4:
		return DEBUG_LEVEL
	case 5:
		return TRACE_LEVEL
	default:
		return WARN_LEVEL
	}
}

type logger struct {
	level  level
	logger *log.Logger
}

type appLogger struct {
	CurrentLevel level
	trace        logger
	debug        logger
	info         logger
	warn         logger
	error        logger
	panic        logger
}

func (logging appLogger) Trace(format string, v ...any) {
	if logging.CurrentLevel >= TRACE_LEVEL {
		logging.trace.logger.Printf(format, v...)
	}
}
func (logging appLogger) Debug(format string, v ...any) {
	if logging.CurrentLevel >= DEBUG_LEVEL {
		logging.debug.logger.Printf(format, v...)
	}
}

func (logging appLogger) Info(format string, v ...any) {
	if logging.CurrentLevel >= INFO_LEVEL {
		logging.info.logger.Printf(format, v...)
	}
}

func (logging appLogger) Warn(format string, v ...any) {
	if logging.CurrentLevel >= WARN_LEVEL {
		logging.warn.logger.Printf(format, v...)
	}
}

func (logging appLogger) Error(format string, v ...any) {
	if logging.CurrentLevel >= ERROR_LEVEL {
		logging.error.logger.Printf(format, v...)
	}
}

func (logging appLogger) Panic(format string, v ...any) {
	if logging.CurrentLevel >= PANIC_LEVEL {
		logging.panic.logger.Panicf(format, v...)
	}
}

var flags = log.Ldate | log.Ltime
var AppLogger = appLogger{
	trace: logger{level: TRACE_LEVEL, logger: log.New(os.Stdout, "TRACE: ", flags)},
	debug: logger{level: DEBUG_LEVEL, logger: log.New(os.Stdout, "DEBUG: ", flags)},
	info:  logger{level: INFO_LEVEL, logger: log.New(os.Stdout, "INFO: ", flags)},
	warn:  logger{level: WARN_LEVEL, logger: log.New(os.Stdout, "WARN: ", flags)},
	error: logger{level: ERROR_LEVEL, logger: log.New(os.Stdout, "ERROR: ", flags)},
	panic: logger{level: PANIC_LEVEL, logger: log.New(os.Stdout, "PANIC: ", flags)},
}
