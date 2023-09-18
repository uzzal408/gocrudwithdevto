package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primeryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"conent"`
	CreatedAt time.Time `json:"created_at"`
}
