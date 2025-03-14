package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     string
	Amount float64
	Status string
}
