package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int `json:"id" gorm:"column(id);primaryKey;autoIncrement;comment:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at;comment:创建时间;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at;comment:创建时间;"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at;comment:删除时间"` // 查询这个字段但是不返回这个字段
}