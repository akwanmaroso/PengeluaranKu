package models

import "time"

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:225;not null" json:"name"`
	Email     string    `gorm:"size:225;not null" json:"email"`
	Password  string    `gorm:"size:225;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
