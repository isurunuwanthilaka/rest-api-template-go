package log

import (
	"log"
	"os"
)

var (
	f      *os.File
	err    error
	logger *log.Logger
)

func init() {
	f, err = os.OpenFile("log.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger = log.New(f, "[Logger]", log.LstdFlags)
}

func GetLogger() *log.Logger {
	return logger
}
