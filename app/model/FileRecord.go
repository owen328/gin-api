package model

import "go_learn/common"

type FileRecord struct {
	BaseModel
	UserId   int    `gorm:"not null"`
	Admin    Admin  `gorm:"foreignKey:UserId"`
	FilePath string `gorm:"type:varchar(255);not null"`
}

func init() {
	common.DB.AutoMigrate(&FileRecord{})
}
