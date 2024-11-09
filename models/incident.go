package models

import "gorm.io/gorm"

type Incident struct {
	gorm.Model
	ProductID   uint
	Description string
	Status      string // e.g., "open", "closed" - possibly enum?
}
