package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TeamID        uint
	Name          string         `gorm:"unique"`
	Email         string         `gorm:"unique"`
	SlackID       *string        `gorm:"unique"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID"`
}
