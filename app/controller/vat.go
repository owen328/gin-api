package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_learn/app/form"
	"go_learn/common"
	"go_learn/service/vies"
	"go_learn/utils"
	"gorm.io/gorm"
	"sync"
)

type IVat interface {
	CheckVat(c *gin.Context)
	CheckMultiVats(c *gin.Context)
}

type Vat struct {
	db *gorm.DB
}

func NewVat() *Vat {
	return &Vat{db: common.DB}
}

func (v Vat) CheckVat(c *gin.Context) {
	var vatForm form.CheckVatForm
	if err := c.ShouldBindJSON(&vatForm); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	result, err := vies.CheckVat(vatForm.VatNumber)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, result)
}

func (v Vat) CheckMultiVats(c *gin.Context) {
	var vatForm form.CheckMultiVatsForm
	if err := c.ShouldBindJSON(&vatForm); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	var wg sync.WaitGroup

	checkRes := make(chan vies.VatResponse, 10)
	for _, vatNumber := range vatForm.VatNumbers {
		fmt.Println(vatNumber)
		wg.Add(1)
		go func(vatNumber string, resChan chan vies.VatResponse) {
			res, _ := vies.CheckVat(vatNumber)
			checkRes <- *res
			wg.Done()
		}(vatNumber, checkRes)
	}
	wg.Wait()
	close(checkRes)
	var resp []vies.VatResponse
	for singleRes := range checkRes {
		resp = append(resp, singleRes)
	}
	utils.Success(c, resp)
	return
}
