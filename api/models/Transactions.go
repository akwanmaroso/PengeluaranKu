package models

import (
	"html"
	"strings"
	"time"
)

type Transaction struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Date        time.Time `gorm:"default:current_timestamp()" json:"date"`
	Amount      uint64    `gorm:"int;size:12; not null" json:"amount"`
	InOut       bool      `gorm:"boolean;not null" json:"in_out"`
	Description string    `gorm:"text" json:"description"`
	Category    Category  `json:"category"`
	CategoryID  uint32    `json:"category_id"`
	Partner     Partner   `json:"partner"`
	PartnerID   uint32    `json:"partner_id"`
	Creator     User      `gorm:"not_null" json:"creator"`
	CreatorID   uint32    `json:"creator_id"`
}

func (p *Transaction) Prepare() {
	p.ID = 0
	p.Date = time.Now()
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Category = Category{}
	p.Partner = Partner{}
	p.Creator = User{}
}
