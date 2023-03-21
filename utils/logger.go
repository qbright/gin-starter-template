package utils

import (
	"gin-starter-template/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var levelMap = map[string]zerolog.Level{
	"trace": zerolog.TraceLevel,
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel := config.GetConfig().GetString("logLevel")
	level := levelMap[logLevel]
	zerolog.SetGlobalLevel(level)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func DLog() *zerolog.Event {
	return log.Debug()
}
func ILog() *zerolog.Event {
	return log.Info()
}

func Elog() *zerolog.Event {
	return log.Error()
}
