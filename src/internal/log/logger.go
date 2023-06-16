package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func init() {
	Logger.Out = os.Stdout

	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
