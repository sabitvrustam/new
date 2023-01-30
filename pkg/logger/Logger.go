package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init() (log *logrus.Logger) {
	log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                   //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = false // remove colors
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout
	return log
}
