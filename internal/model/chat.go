package model

import "time"

type Chat struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`

	Messages []Message `json:"messages" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ChatID    uint      `json:"chat_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
