package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// AperLog : Log entry of aper
var AperLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	AperLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "Aper"})
}

func GetLogger() *logrus.Logger {
	return log
}

// SetLogLevel : set the log level (panic|fatal|error|warn|info|debug|trace)
func SetLogLevel(level logrus.Level) {
	AperLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

// SetReportCaller : Set whether shows the filePath and functionName on loggers
func SetReportCaller(enable bool) {
	AperLog.Infoln("set report call :", enable)
	log.SetReportCaller(enable)
}
