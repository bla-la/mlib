package mlib

import (
	"fmt"
	"time"
)

var KNRM = "\x1B[0m"
var KRED = "\x1B[31m"
var KGRN = "\x1B[32m"
var KYEL = "\x1B[33m"
var KBLU = "\x1B[34m"
var KMAG = "\x1B[35m"
var KCYN = "\x1B[36m"
var KWHT = "\x1B[37m"

func Timestamp() string {
	t0 := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%03d",
		t0.Year(), t0.Month(), t0.Day(), t0.Hour(), t0.Minute(), t0.Second(), t0.Nanosecond()/1000)
}

func _log(severity *string, color *string, format *string, args ...interface{}) {

	msg := fmt.Sprintf(*format, args...)
	fmt.Printf("%s%s %s \"%s\"%s\n",
		*color, Timestamp(),
		*severity,
		msg,
		KNRM)
}

func Error(format string, args ...interface{}) {
	errStr := "ERROR"
	_log(&errStr, &KRED, &format, args...)
}

func Success(format string, args ...interface{}) {
	successStr := "SUCCESS"
	_log(&successStr, &KGRN, &format, args...)
}

func Info(format string, args ...interface{}) {
	infoStr := "INFO"
	_log(&infoStr, &KYEL, &format, args...)
}
