package logger

import (
	"log"

	"github.com/cihub/seelog"
)

func Error(args ...interface{}) {
	seelog.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	seelog.Errorf(format, args...)
}
func Criticalf(format string, args ...interface{}) {
	seelog.Criticalf(format, args...)
}
func Critical(args ...interface{}) {
	seelog.Critical(args...)
}
func Infof(format string, args ...interface{}) {
	seelog.Infof(format, args...)
}
func Info(args ...interface{}) {
	seelog.Info(args...)
}
func Warnf(format string, args ...interface{}) {
	seelog.Warnf(format, args...)
}
func Warn(args ...interface{}) {
	seelog.Warn(args...)
}
func Debugf(format string, args ...interface{}) {
	seelog.Debugf(format, args...)
}
func Debug(args ...interface{}) {
	seelog.Debug(args...)
}

func Tracef(format string, args ...interface{}) {
	seelog.Tracef(format, args...)
}
func Trace(args ...interface{}) {
	seelog.Trace(args...)
}

func Setup(logConfigPath string) {
	logger, err := seelog.LoggerFromConfigAsFile(logConfigPath)
	if err != nil {
		log.Fatalln("err parsing seelog config file:" + err.Error())
		return
	}
	err = seelog.ReplaceLogger(logger)
	if err != nil {
		log.Fatalln("seelog replace logger err:" + err.Error())
		return
	}

	defer seelog.Flush()

}
