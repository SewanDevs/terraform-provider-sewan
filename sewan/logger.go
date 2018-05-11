package sewan

import(
  "io"
  "log"
  "os"
  "strings"
)

func loggerCreate(logFile string) *log.Logger{
	var logger *log.Logger
  var logWriter io.Writer
  var logFolder string
  logFolder = "logs/"

  var logFilePath strings.Builder

  logFilePath.WriteString(logFolder)
  logFilePath.WriteString(logFile)

  _, folder_exists_error := os.Stat(logFolder)
  if folder_exists_error != nil {
    os.Mkdir(logFolder, 0777)
  }

	var _, file_exists_error = os.Stat(logFilePath.String())
	if file_exists_error == nil {
    os.Remove(logFilePath.String())
  }
  os.Create(logFilePath.String())

  logFileObject, logFileErr := os.OpenFile(logFilePath.String(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if logFileErr != nil {
    log.Fatalln("Failed to open log file :", logFileErr)
	}
	logWriter = io.MultiWriter(logFileObject, os.Stdout)
	logger = log.New(logWriter, "Sewan Provider : ", log.Ldate|log.Ltime)

  return logger
}
