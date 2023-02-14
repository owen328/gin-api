package controller

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/form"
	"go_learn/app/model"
	"go_learn/common"
	"go_learn/utils"
	"gorm.io/gorm"
	"os"
	"path"
	"strconv"
	"time"
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
	var request form.UploadSellerReportForm
	if err := c.ShouldBind(&request); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	fileDir := path.Join("uploads", strconv.FormatInt(time.Now().UnixMicro(), 10))
	err = os.MkdirAll(fileDir, 666)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}

	filePath := path.Join(fileDir, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}

	record := &model.TaxCalculateRecord{
		UserId:     c.GetInt("user_id"),
		FilePath:   filePath,
		StartDate:  request.StartDate,
		PeriodType: request.PeriodType,
	}
	tx := s.db.Create(record)
	if tx.Error != nil {
		utils.Fail(c, tx.Error.Error())
		return
	}
	utils.Success(c, nil)
	return
}
