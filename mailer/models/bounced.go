package models

import (
	"time"
)

type Bounced struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Email          string    `json:"email" gorm:"index"`
	Event          string    `json:"event" gorm:"index"`
	EventDateTime  time.Time `json:"event_date_time"`
	Category       string    `json:"category"`
	Reason         string    `json:"reason"`
	CreateDateTime time.Time `json:"create_date_time"`
}
