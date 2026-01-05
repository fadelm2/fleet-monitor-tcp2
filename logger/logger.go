package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func New(level logrus.Level) *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(level)
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &Logger{l}
}
