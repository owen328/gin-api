package route

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/controller"
	"go_learn/app/middleware"
)

func initSellerReportRouter(route *gin.RouterGroup) {

	c := controller.NewSellerReport()
	group := route.Group("/seller-report").Use(middleware.AuthMiddleware())
	{
		group.POST("/upload", c.UploadSellerReport)
	}

}
