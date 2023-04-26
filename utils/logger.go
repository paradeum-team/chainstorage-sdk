package utils

import (
	"chainstorage-sdk/common/plogger"
	"fmt"
	"time"
)

/**
DEBUG
*/
func LogDebug(methodName, msg string) {
	plogger.NewInstance().GetLogger().Debug(fmt.Sprintf("[%s] %s", methodName, msg))
}

/**
INFO
*/
func DevLog(methodName, msg string) {
	plogger.NewInstance().GetLogger().Info(fmt.Sprintf("[%s] %s", methodName, msg))
}

/**
ERROR
*/
func LogError(msg string) {
	plogger.NewInstance().GetLogger().Error(msg)
}

/**
WARNING
*/
func LogWarning(msg string) {
	plogger.NewInstance().GetLogger().Warn(msg)
}

/**
方法耗时
*/
func DevTimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	sprintf := fmt.Sprintf("TIMETRACK: %s took %s", name, elapsed)
	plogger.NewInstance().GetLogger().Debug(sprintf)
}

func DevTimeTrack2(id int, start time.Time, name string) {
	elapsed := time.Since(start)
	sprintf := fmt.Sprintf("%-18s %-8d %s took %s", "timeTrack", id, name, elapsed)
	plogger.NewInstance().GetLogger().Info(sprintf)
}

func DevLog2(id int, methodName, msg string) {
	plogger.NewInstance().GetLogger().Debug(fmt.Sprintf("[%s] %d %s", methodName, id, msg))
}
