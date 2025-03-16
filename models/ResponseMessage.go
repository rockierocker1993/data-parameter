package models

import (
	"gorm.io/gorm"
)

type ResponseMessage struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Code      string
	TitleId   string
	TitleEn   string
	MessageId string
	MessageEn string
	Source    string
	BaseModel
}
