package main

import (
	"fmt"
	"os"
)

func Error(format string, args ...interface{}) {
	fmt.Printf("\033[0;31m"+format+"\033[0m\n", args...)
}

func Warning(format string, args ...interface{}) {
	fmt.Printf("\033[0;33m"+format+"\033[0m\n", args...)
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\033[0;32m"+format+"\033[0m\n", args...)
}

func ErrorAndExit(format string, args ...interface{}) {
	Error(format, args...)
	os.Exit(1)
}
