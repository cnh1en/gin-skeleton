package initialize

import (
	"github.com/cnh1en/lets_go/global"
	"github.com/cnh1en/lets_go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
