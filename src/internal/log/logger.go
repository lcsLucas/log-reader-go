package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func init() {

	file, err := os.OpenFile("/var/log/"+os.Getenv("ENV_PATH")+"/errors.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	Logger.SetOutput(file)

	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
