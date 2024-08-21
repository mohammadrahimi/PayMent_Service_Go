package models

import (
	"time"

	"github.com/google/uuid"
)

type PayMentEntity struct {
	Id              uuid.UUID  `gorm:"primaryKey"` // ";unique"
	Description     string     `gorm:"type:varchar(2000);"`
	CreatedAt       time.Time
	OrderId         uuid.UUID
	UserId          uuid.UUID
	TransactionCode int32      `gorm:"type:int;"`
	TansactionState string     `gorm:"type:varchar(100);"` // NOT NULL
}

 