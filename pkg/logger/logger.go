package logger

import (
	"log"
	"os"
)

type Logger interface{

    ErrorLog(args ...any)
    InfoLog(args ...any)
}


type apiLogger struct{
}

func NewApiLogger() Logger{
    return &apiLogger{}
}
var loggerMap = map[string]*log.Logger{
    "error": log.New(os.Stderr, "ERROR:\t", log.Ldate | log.Ltime),
    "info": log.New(os.Stdout, "INFO:\t", log.Ldate | log.Ltime),
}

func(l *apiLogger)ErrorLog(args ...any){
    loggerMap["error"].Println(args...)
}
func(l *apiLogger)InfoLog(args ...any){
    loggerMap["info"].Println(args...)
}
