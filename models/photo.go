package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" valid:"required~Photo's title is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url" valid:"required~Photo's url is required"`
	UserID   string    `json:"user_id" form:"user_id"`
	User     User      `json:"user"`
	Comments []Comment `json:"Comments" gorm:"constrain:OnDelete:CASCADE"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(p.Title)) == 0 {
		err = errors.New("Photo Title is required")
		return
	}
	if len(strings.TrimSpace(p.PhotoURL)) == 0 {
		err = errors.New("Photo URL is required")
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(p.Title)) == 0 {
		err = errors.New("Photo Title is required")
		return
	}
	if len(strings.TrimSpace(p.PhotoURL)) == 0 {
		err = errors.New("Photo URL is required")
		return
	}

	err = nil
	return
}
