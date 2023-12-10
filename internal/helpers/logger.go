package helpers

import (
	"log"
	"os"
)

type logger struct {
	trace *log.Logger
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	panic *log.Logger
}

func (logging logger) Trace(format string, v ...any) {
	logging.trace.Printf(format, v...)
}
func (logging logger) Debug(format string, v ...any) {
	logging.debug.Printf(format, v...)
}

func (logging logger) Info(format string, v ...any) {
	logging.info.Printf(format, v...)
}

func (logging logger) Warn(format string, v ...any) {
	logging.warn.Printf(format, v...)
}

func (logging logger) Error(format string, v ...any) {
	logging.error.Printf(format, v...)
}

func (logging logger) Panic(format string, v ...any) {
	logging.panic.Panicf(format, v...)
}

var flags = log.Ldate | log.Ltime
var AppLogger = logger{
	trace: log.New(os.Stdout, "TRACE: ", flags),
	debug: log.New(os.Stdout, "DEBUG: ", flags),
	info:  log.New(os.Stdout, "INFO: ", flags),
	warn:  log.New(os.Stdout, "WARN: ", flags),
	error: log.New(os.Stdout, "ERROR: ", flags),
	panic: log.New(os.Stdout, "PANIC: ", flags),
}
