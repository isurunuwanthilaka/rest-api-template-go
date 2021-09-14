package dao

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string
	Age      int8
	GenderId uint
	Gender   Gender `gorm:"foreignKey:GenderId"`
	StatusId uint
	Status   Status `gorm:"foreignKey:StatusId"`
}
