package models

import (
	"errors"
	"strings"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"uniqueIndex;not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~passwrod hast to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if !(u.Age > 16) {
		err = errors.New("Minimun Age is 17")
		return
	}

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(strings.TrimSpace(u.Username)) == 0 {
		err = errors.New("Username is required")
		return
	}

	if len(strings.TrimSpace(u.Email)) == 0 {
		err = errors.New("Email is required")
		return
	}

	if !(govalidator.IsEmail(u.Email)) {
		err = errors.New("Invalid Email Format")
		return
	}

	err = nil
	return
}
