package models

import "time"

type Transaction struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Date        time.Time `gorm:"default:current_timestamp()" json:"date"`
	Amount      uint64    `gorm:"int;size:12; not null" json:"amount"`
	InOut       bool      `gorm:"boolean;not null" json:"int_out"`
	Description string    `gorm:"text" json:"description"`
	Category    Category  `json:"category"`
	CategoryID  uint32    `json:"category_id"`
	Patner      Patner    `json:"patner"`
	PatnerID    uint32    `json:"patner_id"`
	Creator     User      `gorm:"not_null" json:"creator"`
	CreatorID   uint32    `json:"creator_id"`
}
