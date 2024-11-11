package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	TeamID     *uint  `gorm:"uniqueIndex:idx_subscription"`
	UserID     *uint  `gorm:"uniqueIndex:idx_subscription"`
	ProductID  *uint  `gorm:"uniqueIndex:idx_subscription"`
	IncidentID *uint  `gorm:"uniqueIndex:idx_subscription"`
	Type       string `gorm:"uniqueIndex:idx_subscription"`

	Team     Team     `gorm:"foreignKey:TeamID"`
	User     User     `gorm:"foreignKey:UserID"`
	Product  Product  `gorm:"foreignKey:ProductID"`
	Incident Incident `gorm:"foreignKey:IncidentID"`
}
