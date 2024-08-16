package core_utils

import (
	"fmt"
	"log"
	"time"
)

func LogDebug(format string, args ...any) {
	fmt.Printf(time.Now().Format(time.DateTime)+" [DEBUG] "+format+"\r\n", args...)
}

func LogInfo(format string, args ...any) {
	fmt.Printf(time.Now().Format(time.DateTime)+" [INFO] "+format+"\r\n", args...)
}

func LogWarn(format string, args ...any) {
	log.Printf(time.Now().Format(time.DateTime)+" [WARN] "+format+"\r\n", args...)
}

func LogError(format string, args ...any) {
	log.Printf(time.Now().Format(time.DateTime)+" [ERROR] "+format+"\r\n", args...)
}

func LogPanic(format string, args ...any) {
	log.Panicf(time.Now().Format(time.DateTime)+" [PANIC] "+format+"\r\n", args...)
}
