package loghelper

import (
	"fmt"
	"strings"
)

var shouldPrint bool

func SetPrint(input bool) {
	shouldPrint = !input
}

func LogInfo(logVal ...string) {
	joinedInfo := strings.Join(logVal, "-")
	if shouldPrint {

		fmt.Println(joinedInfo)
	}
}

// LogError -  Logs in console and provide Error in logrus
func LogError(logVal ...string) {
	joinedInfo := strings.Join(logVal, "-")
	if shouldPrint {

		fmt.Println(joinedInfo)
	}
}

// LogWarning -  Logs in console and provide warning in logrus
func LogWarning(logVal ...string) {
	joinedInfo := strings.Join(logVal, "-")
	if shouldPrint {

		fmt.Println(joinedInfo)
	}
}
