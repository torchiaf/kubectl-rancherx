package cli

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

var logFileName string
var logLevel uint32

type PlainFormatter struct {
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}

func toggleDebug() error {

	if logFileName != "" {
		logFile := fmt.Sprintf("%s.log", logFileName)

		var f *os.File
		var err error

		if f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			fmt.Println(err)
			return err
		}

		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
	}

	log.SetLevel(log.Level(logLevel))
	log.SetFormatter(&log.TextFormatter{})

	log.Info("Logs enabled")

	return nil
}
