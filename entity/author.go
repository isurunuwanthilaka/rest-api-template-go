package entity

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Id       int64
	Name     string
	Age      int8
	GenderId int8
	Status   int8
}
