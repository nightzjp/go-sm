/*
   @TIME: 2023/6/13 11:39
   @Author: Nightz
   @Site:
   @File: r_user.go
   @SoftWare: GoLand
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"yishi-ai.com/go-sm/handlers"
	"yishi-ai.com/go-sm/middleware"
)

type UserRouter struct{}

var userRouter = new(UserRouter)

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")

	userRouterPub := userRouter // 公共接口
	// 不需要认证访问的接口
	{
		// 用户登录：不需要认证
		userRouterPub.POST("/login", handlers.UserLoginHandler)
		userRouterPub.POST("/register", handlers.UserRegisterHandler)
	}
	userRouterPri := userRouter.Use(middleware.JWTMiddleware()) // 需认证接口
	// 需要认证访问的接口
	{
		userRouterPri.GET("/info", handlers.UserInfoHandler)
	}
}

/*

 */
