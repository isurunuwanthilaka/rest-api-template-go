package dao

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string
}
