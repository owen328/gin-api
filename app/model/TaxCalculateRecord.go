package model

import "go_learn/common"

type TaxCalculateRecord struct {
	BaseModel
	UserId     int    `gorm:"not null"`
	Admin      Admin  `gorm:"foreignKey:UserId"`
	FilePath   string `gorm:"type:varchar(50);not null;comment:销售报表uri"`
	StartDate  string `gorm:"type:date;not null;comment:周期起始"`
	PeriodType int8   `gorm:"type:int(4);default:1;not null;comment:周期类型"`
}

func init() {
	common.DB.AutoMigrate(&TaxCalculateRecord{})
}
