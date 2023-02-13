package controller

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/form"
	"go_learn/common"
	"go_learn/utils"
	"gorm.io/gorm"
)

type ISellerReport interface {
	UploadSellerReport(c *gin.Context)
}

type SellerReport struct {
	db *gorm.DB
}

func NewSellerReport() *SellerReport {
	return &SellerReport{
		db: common.DB,
	}
}

func (s SellerReport) UploadSellerReport(c *gin.Context) {
	var formData form.UploadSellerReportForm
	if err := c.ShouldBindJSON(&formData); err != nil {
		utils.Fail(c, "")
	}
}
