package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"gome/api/security"
	"html"
	"strings"
	"time"
)

// User struct
type User struct {
	ID uint64 `gorm:"primary_id;auto_increment" json:"id"`
	Nickname string `gorm:"size:20;not null;unique" json:"nickname"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:60;not null" json:"password,omitempty"`
	CreateAt time.Time `gorm:"default:current_timestamp()" json:"create_at"`
	UpdateAt time.Time `gorm:"default:current_timestamp()" json:"update_at"`
	Posts []Post `json:"posts,omitempty"`
}

// BeforeSave hash the user password
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare cleans the inputs
func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("nickname is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
	case "login":
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
	default:
		if u.Nickname == "" {
			return errors.New("nickname is require")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
	}
	return nil
}