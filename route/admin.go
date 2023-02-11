package route

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/controller"
	"go_learn/app/middleware"
)


var c *controller.Admin

func init()  {
	c = controller.NewAdmin()
}

func adminRouters(route *gin.RouterGroup)  {
	route.POST("/register", c.Register)
	route.POST("/login", c.Login)
	group := route.Group("").Use(middleware.AuthMiddleware())
	{
		group.GET("/home", c.Home)
	}
	
}