package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID         uint   `json:"user_id"`
	User           User   `json:"user"`
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(u.Name)) == 0 {
		err = errors.New("Social Media Name is required")
		return
	}

	if len(strings.TrimSpace(u.SocialMediaURL)) == 0 {
		err = errors.New("Social Media Url is required")
		return
	}

	err = nil
	return
}
