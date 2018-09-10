package sewansdk

import (
	"io"
	"log"
	"os"
	"strings"
)

// LoggerCreate is a wrapper used only for plugin development debug purpose,
// it must be removed at the end of the develpment cycle, before delivery to prod.
// It creates a logger that write logs to files in sdk-logs/ folder, stored in
// current folder.
func LoggerCreate(logFile string) *log.Logger {
	return loggerCreate(logFile)
}

func loggerCreate(logFile string) *log.Logger {
	logFolder := "sdk-logs/"
	var logFilePath strings.Builder
	logFilePath.WriteString(logFolder)
	logFilePath.WriteString(logFile)
	_, folderExistsError := os.Stat(logFolder)
	if folderExistsError != nil {
		os.Mkdir(logFolder, 0777)
	}
	var _, fileExistsError = os.Stat(logFilePath.String())
	if fileExistsError == nil {
		os.Remove(logFilePath.String())
	}
	os.Create(logFilePath.String())
	logFileObject, logFileErr := os.OpenFile(logFilePath.String(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logFileErr != nil {
		log.Fatalln("Failed to open log file :", logFileErr)
	}
	logWriter := io.MultiWriter(logFileObject)
	return log.New(logWriter, "Sewan Provider : ", log.Ldate|log.Ltime)
}
