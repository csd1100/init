package helpers

import (
	"log"
	"os"
)

type Level int8

const (
	PanicLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func GetLevel(verbosity int) Level {
	switch verbosity {
	case 0:
		return PanicLevel
	case 1:
		return ErrorLevel
	case 2:
		return WarnLevel
	case 3:
		return InfoLevel
	case 4:
		return DebugLevel
	case 5:
		return TraceLevel
	default:
		return WarnLevel
	}
}

type logger struct {
	level  Level
	logger *log.Logger
}

type appLogger struct {
	CurrentLevel Level
	trace        logger
	debug        logger
	info         logger
	warn         logger
	error        logger
	panic        logger
}

func (logging appLogger) Trace(format string, v ...any) {
	if logging.CurrentLevel >= TraceLevel {
		logging.trace.logger.Printf(format, v...)
	}
}
func (logging appLogger) Debug(format string, v ...any) {
	if logging.CurrentLevel >= DebugLevel {
		logging.debug.logger.Printf(format, v...)
	}
}

func (logging appLogger) Info(format string, v ...any) {
	if logging.CurrentLevel >= InfoLevel {
		logging.info.logger.Printf(format, v...)
	}
}

func (logging appLogger) Warn(format string, v ...any) {
	if logging.CurrentLevel >= WarnLevel {
		logging.warn.logger.Printf(format, v...)
	}
}

func (logging appLogger) Error(format string, v ...any) {
	if logging.CurrentLevel >= ErrorLevel {
		logging.error.logger.Printf(format, v...)
	}
}

func (logging appLogger) Panic(format string, v ...any) {
	if logging.CurrentLevel >= PanicLevel {
		logging.panic.logger.Panicf(format, v...)
	}
}

var flags = log.Ldate | log.Ltime
var AppLogger = appLogger{
	trace: logger{level: TraceLevel, logger: log.New(os.Stdout, "TRACE: ", flags)},
	debug: logger{level: DebugLevel, logger: log.New(os.Stdout, "DEBUG: ", flags)},
	info:  logger{level: InfoLevel, logger: log.New(os.Stdout, "INFO: ", flags)},
	warn:  logger{level: WarnLevel, logger: log.New(os.Stdout, "WARN: ", flags)},
	error: logger{level: ErrorLevel, logger: log.New(os.Stdout, "ERROR: ", flags)},
	panic: logger{level: PanicLevel, logger: log.New(os.Stdout, "PANIC: ", flags)},
}
