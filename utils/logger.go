package utils

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() {
	if isInteractive() {
		logrus.SetOutput(os.Stdout)
	} else {
		file, err := os.OpenFile(viper.GetString("logfile_path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.SetOutput(file)
	}
}

func isInteractive() bool {
	return os.Getenv("TERM") != ""
}
