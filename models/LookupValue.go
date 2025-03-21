package models

import (
	"gorm.io/gorm"
)

type LookupValue struct {
	gorm.Model
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Key    string
	Value  string
	TextId string
	TextEn string
	BaseModel
}
