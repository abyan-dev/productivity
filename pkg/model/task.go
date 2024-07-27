package model

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	DueDate     time.Time `gorm:"type:timestamp" json:"due_date"`
	IsComplete  bool      `gorm:"type:boolean" json:"is_complete"`
	UserEmail   string    `gorm:"type:varchar(100);not null" json:"user_email"`
}
