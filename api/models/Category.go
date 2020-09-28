package models

import "time"

type Category struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"text" json:"description"`
	Color       string    `gorm:"char;size:7;default:'#00aabb'" json:"color"`
	Creator     User      `json:"creator"`
	CreatorID   uint32    `gorm:"not null" json:"creator_id"`
	CreatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
