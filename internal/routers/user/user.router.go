package user

import "github.com/gin-gonic/gin"

type UserRouterGroup struct{}

func (r *UserRouterGroup) InitUserRouter(Router *gin.RouterGroup) {

	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.POST("register")
		userRouterPublic.POST("login")
		userRouterPublic.POST("otp")
	}

	userRouterPrivate := Router.Group("user")
	{
		userRouterPrivate.GET("/:id")
	}

}
