package models

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey,autoIncrement" json:"id"`
	Title     string    `gorm:"type:string" json:"title"`
	Body      string    `gorm:"type:longText" json:"body"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
