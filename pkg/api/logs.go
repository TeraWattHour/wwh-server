package api

import (
	"log"
	"os"
)

var InternalLogger *log.Logger

func SetupLoggers() error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	InternalLogger = log.New(file, "INTERNAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}
