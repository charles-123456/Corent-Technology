package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
)
import "gopkg.in/natefinch/lumberjack.v2"

var (
	logObj *myLogger
)

type CustomLog interface {
	Write(v ...interface{})
	PanicLogger()
}

const Ldate = log.Ldate
const Ltime = log.Ltime
const Lshortfile = log.Lshortfile
const maxfilesize = 25

//LoggerInstance contains the logger instance for info,error,trace and warning logs
type myLogger struct {
	traceLog   *log.Logger
	errLog     *log.Logger
	warnLog    *log.Logger
	infoLog    *log.Logger
	panicLog   *log.Logger
	statusLog  *log.Logger
	userLog    *log.Logger
	printTrace bool
	printInfo  bool
	printError bool
	printWarn  bool
	loggerPath string
}

type customLogger struct {
	logger *log.Logger
}

//New method creates a new Instance of logger for given module
//userpath specifies the log file location.if userpath is empty then log files will be created in Current working directory.
func NewLogger(logPath string) error {
	if logPath == "" {
		return errors.New("logPath is empty")
	}
	dir, _ := filepath.Split(logPath)
	err := os.MkdirAll(dir, 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new log: %s", err)
	}
	logObj = new(myLogger)
	logObj.loggerPath = logPath
	logObj.infoLog = log.New(&lumberjack.Logger{Filename: logPath + "/info.log", MaxSize: maxfilesize}, "", 0)
	logObj.errLog = log.New(&lumberjack.Logger{Filename: logPath + "/error.log", MaxSize: maxfilesize}, "", 0)
	logObj.warnLog = log.New(&lumberjack.Logger{Filename: logPath + "/warn.log", MaxSize: maxfilesize}, "", log.Ldate|log.Ltime|log.Lshortfile)
	logObj.traceLog = log.New(&lumberjack.Logger{Filename: logPath + "/trace.log", MaxSize: maxfilesize}, "", log.Ldate|log.Ltime|log.Lshortfile)
	logObj.statusLog = log.New(&lumberjack.Logger{Filename: logPath + "/status.log", MaxSize: maxfilesize}, "", 0)
	logObj.panicLog = log.New(&lumberjack.Logger{Filename: logPath + "/panic.log", MaxSize: maxfilesize}, "", log.Ldate|log.Ltime|log.Lshortfile)
	logObj.printTrace = true
	logObj.printError = true
	logObj.printInfo = true
	logObj.printWarn = true
	return nil
}

//Trace method prints log in trace.log
func Trace(v ...interface{}) {
	if logObj.printTrace {
		logObj.trace(v...)
	}
}

//Error method prints log in error.log
func Error(v ...interface{}) {
	if logObj.printError {
		logObj.errors(v...)
	}
}

//Info method prints log in info.log
func Info(v ...interface{}) {
	if logObj.printInfo {
		logObj.info(v...)
	}
}

//Warn method prints log in warn.log
func Warn(v ...interface{}) {
	if logObj.printWarn {
		logObj.warn(v...)
	}
}

func Status(v ...interface{}) {
	err := logObj.statusLog.Output(3, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("logger output :", err)
	}
}

func NewCustomLogger(filename, prefix string, lineFlag int) (*customLogger, error) {
	dir, _ := filepath.Split(filename)
	err := os.MkdirAll(dir, 0744)
	if err != nil {
		return nil, fmt.Errorf("can't make directories for new logfile: %s", err)
	}
	cLogger := new(customLogger)
	cLogger.logger = log.New(&lumberjack.Logger{Filename: filename, MaxSize: 25}, prefix, lineFlag)
	return cLogger, nil
}

func Panic(v ...interface{}) {
	err := logObj.panicLog.Output(3, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("logger output :", err)
	}
}

func (logObj *myLogger) trace(v ...interface{}) {
	err := logObj.traceLog.Output(3, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("logger output :", err)
	}

}

func (logObj *myLogger) errors(v ...interface{}) {
	err := logObj.errLog.Output(3, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("logger output :", err)
	}
}

func (logObj *myLogger) info(v ...interface{}) {
	err := logObj.infoLog.Output(3, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("logger output :", err)
	}
}

func (logObj *myLogger) warn(v ...interface{}) {
	err := logObj.warnLog.Output(3, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("logger output :", err)
	}
}

func DeleteOldLogs() {
	list, err := getFileList(logObj.loggerPath)
	if err == nil {
		for _, val := range list {
			_, filename := filepath.Split(val)
			if !(filename == "info.log" || filename == "error.log" || filename == "warn.log" || filename == "trace.log") {
				err := os.Remove(val)
				if err != nil {
					Error("Delete old logs:", err)
				}
			}
		}
	}
}

//SetLevels method parameter 'levels' specifies the log levels that has to be recorded.
//by default it allows 'Error,Info,Trace and Warning'
func SetLevels(trace, info, warn, errors bool) {
	logObj.printTrace = trace
	logObj.printInfo = info
	logObj.printError = errors
	logObj.printWarn = warn
}

//func Close() {
//	logObj.infoIOWriter.Close()
//	logObj.errorIOWriter.Close()
//	logObj.traceIOWriter.Close()
//	logObj.warnIOWriter.Close()
//	logObj.panicIOWriter.Close()
//}

func PanicLogger() {
	if r := recover(); r != nil {
		Panic("Panic Recovery:", r)
		Panic("----------------------------------------------------------------------------------")
		Panic(string(debug.Stack()))
		Panic("----------------------------------------------------------------------------------")
	}
}

func GetLogger() *myLogger {
	return logObj
}

func getFileList(srcPath string) ([]string, error) {
	var streamList []string
	var DirWalker filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		streamList = append(streamList, path)
		return nil
	}
	err := filepath.Walk(srcPath, DirWalker)
	if err != nil {
		return nil, err
	}
	if len(streamList) < 2 {
		return nil, errors.New("folder is empty")
	}
	streamList = streamList[1:]
	return streamList, nil
}

func (cLogger *customLogger) Write(v ...interface{}) {
	err := cLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println("customLogger :", err)
	}
}

func (cLogger *customLogger) PanicLogger() {
	if r := recover(); r != nil {
		err := cLogger.logger.Output(2, fmt.Sprintln("Panic Recovery:", r))
		if err != nil {
			fmt.Println("customLogger :", err)
			return
		}
		cLogger.logger.Output(2, fmt.Sprintln("----------------------------------------------------------------------------------"))
		cLogger.logger.Output(2, fmt.Sprintln(string(debug.Stack())))
		cLogger.logger.Output(2, fmt.Sprintln("----------------------------------------------------------------------------------"))
	}
}
