package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name      string     `gorm:"unique"`
	Incidents []Incident `gorm:"foreignKey:ProductID"`
}
