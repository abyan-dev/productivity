package model

import (
	"time"
)

type PomodoroSession struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StartTime time.Time `gorm:"type:timestamp;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:timestamp;not null" json:"end_time"`
	UserEmail string    `gorm:"type:varchar(100);not null" json:"user_email"`
}
