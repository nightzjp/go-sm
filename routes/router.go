/*
   @TIME: 2023/6/13 11:18
   @Author: Nightz
   @Site:
   @File: router.go
   @SoftWare: GoLand
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	PrivateGroup := Router.Group("/api/v1")
	{
		userRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}

/*

 */
