package utils

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() {
	log := logrus.New()

	if isInteractive() {
		log.Out = os.Stdout
	} else {
		file, err := os.OpenFile(viper.GetString("logfile_path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.Out = file
	}
}

func isInteractive() bool {
	return os.Getenv("TERM") != ""
}
