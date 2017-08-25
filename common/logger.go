package common

import (
	"fmt"
	"log"
	"os"
)

const (
	LogFilePath           = ""
	FileFlag              = os.O_APPEND | os.O_RDWR | os.O_CREATE
	LogFlag               = log.Ldate | log.Ltime | log.Llongfile
	CarriageReturnNewLine = "\r\n"
	FileLogMode           = 0
	ConsoleLogMode        = 0777
)

type Logger interface {
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type LoggerHelper struct {
	logFile *os.File
}

func (h *LoggerHelper) getLoggerFile(path string, perm os.FileMode) (logFile *os.File, err error) {
	if logFile, err = os.OpenFile(path, FileFlag, perm); err != nil {
		fmt.Printf("Load log file failed, error: %s", err.Error())
		return
	}

	return
}

func (h *LoggerHelper) getLogger(prefix string, flag int) (logger *log.Logger, err error) {
	if h.logFile, err = h.getLoggerFile(LogFilePath, FileLogMode); err != nil {
		return
	}

	logger = log.New(h.logFile, CarriageReturnNewLine, FileFlag)

	return
}

var baseLoggerHelper = LoggerHelper{}

type FileLogger struct {
}

func (f *FileLogger) Info(format string, v ...interface{}) {
	if logger, err := baseLoggerHelper.getLogger(CarriageReturnNewLine, FileLogMode); err != nil {
		defer baseLoggerHelper.logFile.Close()
		logger.Printf(format, v...)
	}
}

func (f *FileLogger) Error(format string, v ...interface{}) {
	if logger, err := baseLoggerHelper.getLogger(CarriageReturnNewLine, FileLogMode); err != nil {
		defer baseLoggerHelper.logFile.Close()
		logger.Fatalf(format, v...)
	}
}

type ConsoleLogger struct {
}

func (c *ConsoleLogger) Info(format string, v ...interface{}) {
	if logger, err := baseLoggerHelper.getLogger(CarriageReturnNewLine, ConsoleLogMode); err != nil {
		defer baseLoggerHelper.logFile.Close()
		logger.Printf(format, v...)
	}
}

func (c *ConsoleLogger) Error(format string, v ...interface{}) {
	if logger, err := baseLoggerHelper.getLogger(CarriageReturnNewLine, ConsoleLogMode); err != nil {
		defer baseLoggerHelper.logFile.Close()
		logger.Fatalf(format, v...)
	}
}

var LoggerType = struct {
	File    int
	Console int
}{
	0,
	1,
}

type LoggerFactory interface {
	Create() (logger *Logger)
}

type FileLoggerFactory struct {
}

func (*FileLoggerFactory) Create() (logger *FileLogger) {
	return &FileLogger{}
}

type ConsoleLoggerFactory struct {
}

func (*ConsoleLoggerFactory) Create() (logger *ConsoleLogger) {
	return &ConsoleLogger{}
}
