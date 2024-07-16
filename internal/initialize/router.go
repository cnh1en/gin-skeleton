package initialize

import (
	"github.com/cnh1en/lets_go/global"
	"github.com/cnh1en/lets_go/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine

	if global.Server.Mode == "local" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("checkStatus")
	}

	{
		userRouter.InitUserRouter(MainGroup)
	}

	return r
}
