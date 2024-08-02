package app

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func setLogLevel(level string) {
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	logrus.SetOutput(os.Stdout)

}
