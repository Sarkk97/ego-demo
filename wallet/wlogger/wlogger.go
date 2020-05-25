package wlogger

import (
	"fmt"
	"log"
	"os"
)

const (
	//INFO is  log severity type info
	INFO = "INFO"
	//ERROR is log severity type error
	ERROR = "ERROR"
)

//Logger is a Logger
type Logger interface {
	log(message string, severtiy string)
	Error(message string)
	Info(message string)
}

//FileLogger is a Logger implementation that logs to a file
type FileLogger struct {
	logger *log.Logger
}

//Log prints a log message
func (fileLogger FileLogger) log(message string, severity string) {
	f, err := os.OpenFile("logs/wallet.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}

	defer f.Close()

	fileLogger.logger = log.New(f, fmt.Sprintf("wallet.%s ", severity), log.LstdFlags)
	fileLogger.logger.Println(message)
}

//Error logs error messages
func (fileLogger FileLogger) Error(message string) {
	fileLogger.log(message, ERROR)
}

//Info logs infomational messages
func (fileLogger FileLogger) Info(message string) {
	fileLogger.log(message, INFO)
}

//NewLogger instantiates a Logger
func NewLogger() Logger {
	return &FileLogger{logger: nil}
}
