package models

import (
	"github.com/asaskevich/govalidator"
	"go-pos/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string        `json:"name" gorm:"not null" valid:"required~Name is required"`
	Username    string        `json:"username" gorm:"unique;not null;uniqueIndex" valid:"required~Username is required"`
	Password    string        `gorm:"not null" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters" json:"password"`
	Transaction []Transaction `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate

		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
