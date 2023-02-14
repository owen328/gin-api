package route

import (
	"github.com/gin-gonic/gin"
	"go_learn/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.LogRequestMiddleware())

	router.StaticFS("uploads", gin.Dir("uploads", false))
	initGroupRoute(router)
	return router
}

func initGroupRoute(router *gin.Engine) {
	apiGroup := router.Group("/api")
	initAdminRouter(apiGroup)
	initSellerReportRouter(apiGroup)
	initVatRouter(apiGroup)
}
