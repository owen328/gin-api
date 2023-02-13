package route

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/controller"
	"go_learn/app/middleware"
)

func initAdminRouter(route *gin.RouterGroup) {

	c := controller.NewAdmin()
	routerGroup := route.Group("/admin")
	routerGroup.POST("/register", c.Register)
	routerGroup.POST("/login", c.Login)
	group := routerGroup.Use(middleware.AuthMiddleware())
	{
		group.GET("/home", c.Home)
	}

}
