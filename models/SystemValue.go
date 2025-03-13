package models

import (
	"gorm.io/gorm"
)

type SystemValue struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Module    string
	Key       string
	Value     string
	IsEncrypt bool
	BaseModel
}
