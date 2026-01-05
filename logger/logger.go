package logger

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	// pastikan folder logs ada
	logDir := "logs"
	_ = os.MkdirAll(logDir, 0755)

	logFile := filepath.Join(logDir, "app.log")

	file, err := os.OpenFile(logFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		panic(err)
	}

	Log.SetOutput(file)

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Log.SetLevel(logrus.InfoLevel)
}
