package model

import "go_learn/common"

type Admin struct {
	BaseModel
	Username   string `json:"username" gorm:"type:varchar(50);column(username);not null;unique;comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(200);column(password);not null;comment:密码"`
	Mobile string `json:"mobile" gorm:"type:varchar(20);column(mobile);not null;unique;comment:手机"`
}

func init() {
	common.DB.AutoMigrate(&Admin{})
}