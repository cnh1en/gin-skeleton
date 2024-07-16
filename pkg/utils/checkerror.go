package utils

import (
	"github.com/cnh1en/lets_go/global"
	"go.uber.org/zap"
)

func CheckErrorPanic(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg, zap.Error(err))
		panic(msg)
	}
}
