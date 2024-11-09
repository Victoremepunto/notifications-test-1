package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name          string         `gorm:"unique"`
	Users         []User         `gorm:"foreignKey:TeamID"`
	Subscriptions []Subscription `gorm:"foreignKey:TeamID"`
}
