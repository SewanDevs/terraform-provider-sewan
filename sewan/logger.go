package sewan

import(
  "io"
  "log"
  "os"
)

func loggerCreate(logFilePath string) *log.Logger{
	var logger *log.Logger
  var logWriter io.Writer

  logFile, logFileErr := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logFileErr != nil {
		log.Fatalln("Failed to open log file :", logFileErr)
	}
	logWriter = io.MultiWriter(logFile, os.Stdout)
	logger = log.New(logWriter, "Sewan Provider : ", log.Ldate|log.Ltime)

  return logger
}
