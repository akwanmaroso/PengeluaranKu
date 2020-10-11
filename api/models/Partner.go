package models

import "time"

type Partner struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:60;not null" json:"name"`
	Description string    `gorm:"text;not null" json:"description"`
	Creator     User      `json:"creator"`
	CreatorID   uint32    `gorm:"not null" json:"creator_id"`
	CreatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
