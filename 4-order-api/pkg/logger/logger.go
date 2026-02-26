package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)

	// JSON формат
	l.SetFormatter(&logrus.JSONFormatter{})

	// уровень можно менять через env, но пока так
	l.SetLevel(logrus.InfoLevel)

	return l
}
