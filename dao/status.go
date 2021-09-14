package dao

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name string
}
