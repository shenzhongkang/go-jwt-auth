package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

// Post model
type Post struct {
	ID uint64 `gorm:"primary_id:auto_increment" json:"id"`
	Title string `gorm:"size:30;not null;unique" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Author User `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID uint64 `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

// Prepare cleans the inputs
func (p *Post) Prepare()  {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
}

// Validate validates the inputs
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	if p.AuthorID < 1 {
		return errors.New("author is required")
	}

	return nil
}