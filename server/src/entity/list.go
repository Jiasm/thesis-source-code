package entity

import "time"

type List struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreateTime time.Time `json:"create_time"`
}

