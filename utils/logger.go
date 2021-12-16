package utils

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var LogFilePath = ""

func InitializeLogging() error {
	if LogFilePath == "" {
		return errors.New("please provide a path for logger")
	}
	var file, err = os.OpenFile(LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(file)
	return nil
}

func LogInfo(message string, component string) {
	log.WithField("component", component).Info(message)
}
func LogError(message string, component string) {
	log.WithField("component", component).Error(message)

}

func LogFatal(message string, component string) {
	log.WithField("component", component).Fatal(message)
	panic(1)
}
