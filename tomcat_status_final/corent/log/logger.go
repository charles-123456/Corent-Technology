package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type CustomLog interface {
	Write(v ...interface{})
	PanicLogger()
}

const (
	maxfilesize = 25
	LevelPanic  = 0
	LevelError  = 1
	LevelWarn   = 2
	LevelInfo   = 3
	LevelTrace  = 4
)

type Level int8

var levelMap = map[Level]string{
	LevelPanic: "panic",
	LevelError: "error",
	LevelWarn:  "warn",
	LevelInfo:  "info",
	LevelTrace: "trace",
}

//LoggerInstance contains the logger instance for info,error,trace and warning logs
var (
	traceLogger *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	panicLogger *log.Logger
	logLevel    Level
	loggerPath  string
	currentPath string
)

type customLogger struct {
	logger *log.Logger
}

type Config struct {
	MaxFileSize  int
	AppendPrefix bool
	Flags        bool
}

func init() {
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	loggerPath = currentPath
	setLoggers(currentPath)
	logLevel = LevelTrace
}

func SetLogPath(logPath string) error {
	if logPath == "" {
		return fmt.Errorf("Given logPath is empty")
	}
	//	logPath = logPath + "/log"
	err := os.MkdirAll(logPath, 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new log: %s", err)
	}
	loggerPath = logPath
	setLoggers(logPath)
	return nil
}

func setLoggers(path string) {
	infoLogger = log.New(&lumberjack.Logger{Filename: path + "/info.log", MaxSize: maxfilesize}, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(&lumberjack.Logger{Filename: path + "/error.log", MaxSize: maxfilesize}, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(&lumberjack.Logger{Filename: path + "/warn.log", MaxSize: maxfilesize}, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	traceLogger = log.New(&lumberjack.Logger{Filename: path + "/trace.log", MaxSize: maxfilesize}, "TRACE: ", log.Ldate|log.Ltime)
	panicLogger = log.New(&lumberjack.Logger{Filename: path + "/panic.log", MaxSize: maxfilesize}, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogPath() string {
	return loggerPath
}

//Trace method prints log in trace.log
func Trace(v ...interface{}) {
	if logLevel >= LevelTrace {
		printTrace(v...)
	}
}

func Tracef(format string, a ...interface{}) {
	if logLevel >= LevelTrace {
		printTrace(fmt.Sprintf(format, a...))
	}
}

func getLogger(filename, prefix string, filesize, flag int) *log.Logger {
	return log.New(&lumberjack.Logger{Filename: loggerPath + "/" + filename + ".log", MaxSize: filesize}, prefix, flag)
}

func alter(logLevel Level, conf *Config) {
	prefix := ""
	var filesize = maxfilesize
	if conf.MaxFileSize > 0 {
		filesize = conf.MaxFileSize
	}
	filename := levelMap[logLevel]
	if conf.AppendPrefix {
		prefix = fmt.Sprintf("%s: ", strings.ToUpper(filename))
	}
	var flag int
	if conf.Flags {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}
	switch logLevel {
	case LevelPanic:
		panicLogger = getLogger(filename, prefix, filesize, flag)
	case LevelInfo:
		infoLogger = getLogger(filename, prefix, filesize, flag)
	case LevelError:
		errorLogger = getLogger(filename, prefix, filesize, flag)
	case LevelTrace:
		traceLogger = getLogger(filename, prefix, filesize, flag)
	case LevelWarn:
		warnLogger = getLogger(filename, prefix, filesize, flag)
	}

}

func Alter(conf *Config, levels ...Level) {
	for _, level := range levels {
		alter(level, conf)
	}
}

//Error method prints log in error.log
func Error(v ...interface{}) {
	if logLevel >= LevelError {
		printError(v...)
	}
}

func Errorf(format string, a ...interface{}) {
	if logLevel >= LevelError {
		printError(fmt.Sprintf(format, a...))
	}
}

//Info method prints log in info.log
func Info(v ...interface{}) {
	if logLevel >= LevelInfo {
		printInfo(v...)
	}
}

func Infof(format string, a ...interface{}) {
	if logLevel >= LevelInfo {
		printInfo(fmt.Sprintf(format, a...))
	}
}

//Warn method prints log in warn.log
func Warn(v ...interface{}) {
	if logLevel >= LevelWarn {
		printWarn(v...)
	}
}

func Warnf(format string, a ...interface{}) {
	if logLevel >= LevelWarn {
		printWarn(fmt.Sprintf(format, a...))
	}
}

func NewCustomLogger(filename, prefix string, hasFlags bool) (*customLogger, error) {
	dir, _ := filepath.Split(filename)
	err := os.MkdirAll(dir, 0744)
	if err != nil {
		return nil, fmt.Errorf("can't make directories for new logfile: %s", err)
	}
	var flag int
	if hasFlags {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}
	cLogger := new(customLogger)
	cLogger.logger = log.New(&lumberjack.Logger{Filename: filename, MaxSize: 25}, prefix, flag)
	return cLogger, nil
}

func Panic(v ...interface{}) {
	if err := panicLogger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Fprintln(os.Stderr, "print panic log :", err)
	}
}

func Panicf(format string, a ...interface{}) {
	if err := panicLogger.Output(2, fmt.Sprintf(format, a...)); err != nil {
		fmt.Fprintln(os.Stderr, "print panic log :", err)
	}
}

func printTrace(v ...interface{}) {
	var funcName, fileName, line string
	var callerSlice = make([]uintptr, 10)
	runtime.Callers(3, callerSlice)
	rf := runtime.CallersFrames(callerSlice)
	frame, ok := rf.Next()
	if !ok {
		fmt.Fprintln(os.Stderr, "Could not find func name")
		fileName = "???"
		funcName = "?"
		line = "?"
	} else {
		_, funcNameWithPackage := filepath.Split(frame.Function)
		lastDot := strings.LastIndex(funcNameWithPackage, ".")
		funcName = funcNameWithPackage[lastDot+1:]
		fileName = GetShortFileName(frame.File)
		fileName = strings.Replace(fileName, "/", ".", -1)
		line = fmt.Sprint(frame.Line)
	}
	traceLogger.Print(fmt.Sprint(fmt.Sprintf("%s:%s-%s()", fileName, line, funcName), " ", fmt.Sprintln(v...)))
	//	if err := traceLogger.Output(3, fmt.Sprintln(funcName, fmt.Sprintln(v...))); err != nil {
	//		fmt.Fprintln(os.Stderr, "print trace log :", err)
	//	}
}

func printError(v ...interface{}) {
	if err := errorLogger.Output(3, fmt.Sprintln(v...)); err != nil {
		fmt.Fprintln(os.Stderr, "print error log :", err)
	}
}

func printInfo(v ...interface{}) {
	if err := infoLogger.Output(3, fmt.Sprintln(v...)); err != nil {
		fmt.Fprintln(os.Stderr, "print info log :", err)
	}
}

func printWarn(v ...interface{}) {
	if err := warnLogger.Output(3, fmt.Sprintln(v...)); err != nil {
		fmt.Fprintln(os.Stderr, "print warn log :", err)
	}
}

func DeleteOldLogs() {
	list, err := getFileList(loggerPath)
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
func SetLevels(level Level) {
	logLevel = level
}

//func Close() {
//	logObj.infoIOWriter.Close()
//	logObj.errorIOWriter.Close()
//	logObj.traceIOWriter.Close()
//	logObj.warnIOWriter.Close()
//	logObj.panicIOWriter.Close()
//}

// func PanicLogger() {
// 	if r := recover(); r != nil {
// 		Panic("Panic Recovery:", r)
// 		Panic("----------------------------------------------------------------------------------")
// 		Panic(string(debug.Stack()))
// 		Panic("----------------------------------------------------------------------------------")
// 	}
// }

func (l *customLogger) PanicLogger() {
	if r := recover(); r != nil {
		l.Write("Panic Recovery:", r)
		l.Write("----------------------------------------------------------------------------------")
		l.Write(string(debug.Stack()))
		l.Write("----------------------------------------------------------------------------------")
	}
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
	if err := cLogger.logger.Output(2, fmt.Sprintln(v...)); err != nil {
		fmt.Fprintln(os.Stderr, "customLogger :", err)
	}
}

func GetShortFileName(fileNameWithExt string) string {
	filename := fileNameWithExt[strings.LastIndex(fileNameWithExt, "src")+4:]
	var result = filename
	if len(filename) > 25 {
		packageNames := strings.Split(filename, "/")
		for i := 0; i < len(packageNames)-1; i++ {
			packageNames[i] = packageNames[i][:1]
		}
		result = strings.Join(packageNames, "/")
	}
	return strings.Replace(result, "/", ".", -1)
}
