package models

import (
	"html"
	"strings"
	"time"

	helper "github.com/akwanmaroso/PengeluaranKu/api/helpers"
)

// User models
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:225;not null" json:"name"`
	Email     string    `gorm:"size:225;not null;unique" json:"email"`
	Password  string    `gorm:"size:225;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// BeforeSave use to change password to hashedPassword before save to database
func (u *User) BeforeSave() error {
	hashedPassword, err := helper.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
