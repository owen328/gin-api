package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(50);column(username);not null;unique;comment:用户名"`
	Password string `gorm:"type:varchar(200);column(password);not null;comment:密码"`
}