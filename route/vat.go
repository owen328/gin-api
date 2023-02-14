package route

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/controller"
	"go_learn/app/middleware"
)

func initVatRouter(route *gin.RouterGroup) {
	c := controller.NewVat()
	routerGroup := route.Group("/vat").Use(middleware.AuthMiddleware())
	{
		routerGroup.POST("/check", c.CheckVat)
		routerGroup.POST("/multi-check", c.CheckMultiVats)
	}

}
