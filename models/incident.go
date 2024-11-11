package models

import "gorm.io/gorm"

type Incident struct {
	gorm.Model
	ProductID   uint
	Name        string `gorm:"unique"`
	Description string
	Status      string // e.g., "open", "closed" - possibly enum?
}
