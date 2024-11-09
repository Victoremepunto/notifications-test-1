package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model

	SubscriptionID uint
	SentAt         *time.Time
	Status         string // e.g., "sent", "failed"
	Content        string

	Subscription Subscription `gorm:"foreignKey:SubscriptionID"`
}
