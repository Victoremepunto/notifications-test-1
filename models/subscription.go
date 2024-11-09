package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	TeamID              *uint  `gorm:"index"`
	UserID              *uint  `gorm:"index"`
	ProductID           uint   `gorm:"index"`
	IncidentID          uint   `gorm:"index"`
	NotificationChannel string `gorm:"type:enum('email', 'slack')"`

	Team     Team     `gorm:"foreignKey:TeamID"`
	User     User     `gorm:"foreignKey:UserID"`
	Product  Product  `gorm:"foreignKey:ProductID"`
	Incident Incident `gorm:"foreignKey:IncidentID"`

	// Unique constraint on combination of team/user, product, incident, and channel
	//gorm:"uniqueIndex:idx_subscription_unique,unique"`
}
