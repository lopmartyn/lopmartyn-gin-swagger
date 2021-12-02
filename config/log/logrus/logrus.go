package logrus

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	//Set log output to standard output (default output is stderr, standard error)
	//Log message output can be any io.writer type
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// log.InfoLevel to show all logs.
	// log.WarnLevel to show both of warning and error.
	// log.ErrorLevel to only error.
	log.SetLevel(log.InfoLevel)

}
