package logger

import "fmt"

// Debug - shorthand debug.
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Info - shorthand info.
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Warn - shorthand warn.
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Error - shorthand error.
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Fatal - shorthand fatal.
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

// Panic - shorthand panic.
func Panic(args ...interface{}) {
	Log.Panic(args...)
}

// Debugf - shorthand debug using format.
func Debugf(format string, args ...interface{}) {
	Debug(fmt.Sprintf(format, args...))
}

// Infof - shorthand info using format.
func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}

// Warnf - shorthand warn using format.
func Warnf(format string, args ...interface{}) {
	Warn(fmt.Sprintf(format, args...))
}

// Errorf - shorthand error using format.
func Errorf(format string, args ...interface{}) {
	Error(fmt.Sprintf(format, args...))
}

// Fatalf - shorthand fatal using format.
func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}

// Panicf - shorthand panic using format.
func Panicf(format string, args ...interface{}) {
	Panic(fmt.Sprintf(format, args...))
}
